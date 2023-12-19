package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lingesh-sk/spotify_usercase/model"
	"github.com/lingesh-sk/spotify_usercase/service"
)

type Routes struct {
	TrackService *service.TrackService
}

func NewRoutes(trackService *service.TrackService) *Routes {
	return &Routes{
		TrackService: trackService,
	}
}

func (r *Routes) RegisterRoutes(router *gin.Engine) {
	router.GET("/track/:isrc", r.GetTrackDetailsByISRC)
	router.GET("/track/artist/:artistName", r.SearchTracksByArtistName)
	router.POST("/track", r.CreateTrack)
	router.PUT("/track/:isrc", r.UpdateTrack)
}

// GetTrackDetailsByISRC
func (r *Routes) GetTrackDetailsByISRC(c *gin.Context) {
	isrc := c.Param("isrc")
	trackDetails, err := r.TrackService.GetTrackDetailsByISRC(isrc)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, trackDetails)
}

// SearchTracksByArtistName
func (r *Routes) SearchTracksByArtistName(c *gin.Context) {
	artistName := c.Param("artistName")
	trackDetailsList, err := r.TrackService.SearchTracksByArtistName(artistName)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, trackDetailsList)
}

// CreateTrack 
func (r *Routes) CreateTrack(c *gin.Context) {
	var trackDetails model.TrackDetails
	if err := c.ShouldBindJSON(&trackDetails); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if ISRC code already exists
	existingTrack, err := r.TrackService.GetTrackDetailsByISRC(trackDetails.ISRC)
	if err == nil {
		c.JSON(200, existingTrack)
		return
	}

	// If track not found in db or Spotify, then create a new record
	createdTrack, err := r.TrackService.GetOrCreateTrackDetails(trackDetails.ISRC)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Track record created successfully", "track": createdTrack})
}

// UpdateTrack 
func (r *Routes) UpdateTrack(c *gin.Context) {
	isrc := c.Param("isrc")
	var updatedTrackDetails model.TrackDetails
	if err := c.ShouldBindJSON(&updatedTrackDetails); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	trackDetails, err := r.TrackService.UpdateTrack(isrc, &updatedTrackDetails)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, trackDetails)
}
