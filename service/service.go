package service

import (
	"fmt"
	"github.com/lingesh-sk/spotify_usercase/dao"
	"github.com/lingesh-sk/spotify_usercase/model"
	"github.com/zmb3/spotify"
)

type SpotifyService struct {
	SpotifyClient *spotify.Client
}

func NewSpotifyService(client *spotify.Client) *SpotifyService {
	return &SpotifyService{SpotifyClient: client}
}

// SearchTrackByISRC using the Spotify API
func (s *SpotifyService) SearchTrackByISRC(isrc string) (*spotify.FullTrack, error) {
	query := fmt.Sprintf("isrc:%s", isrc)
	searchResult, err := s.SpotifyClient.Search(query, spotify.SearchTypeTrack)
	if err != nil || len(searchResult.Tracks.Tracks) == 0 {
		return nil, fmt.Errorf("track not found")
	}
	return &searchResult.Tracks.Tracks[0], nil
}

type TrackService struct {
	DatabaseAccessor *dao.DatabaseAccessor
	SpotifyService   *SpotifyService
}

func NewTrackService(dbAccessor *dao.DatabaseAccessor, spotifyService *SpotifyService) *TrackService {
	return &TrackService{DatabaseAccessor: dbAccessor, SpotifyService: spotifyService}
}

// @Summary Get track details by ISRC
// @Description Get track details from the database or Spotify by ISRC code
// @ID get-track-by-isrc
// @Produce json
// @Param isrc path string true "ISRC code of the track"
// @Success 200 {object} model.TrackDetails
// @Failure 404 {object} map[string]interface{} "Track not found"
// @Router /track/{isrc} [get]
// GetTrackDetailsByISRC
func (ts *TrackService) GetTrackDetailsByISRC(isrc string) (*model.TrackDetails, error) {
	existingTrack, err := ts.DatabaseAccessor.GetTrackByISRC(isrc)

	if err == nil {
		trackDetails := &model.TrackDetails{
			ISRC:         existingTrack.ISRC,
			Title:        existingTrack.Title,
			ArtistName:   existingTrack.ArtistName,
			SpotifyImage: existingTrack.SpotifyImage,
		}
		return trackDetails, nil
	}

	// If Track not found in the db then search in Spotify
	track, err := ts.SpotifyService.SearchTrackByISRC(isrc)
	if err != nil {
		return nil, fmt.Errorf("track not found")
	}

	// Extracting details from  Spotify search
	trackDetails := &model.TrackDetails{
		ISRC:         isrc,
		Title:        track.Name,
		ArtistName:   track.Artists[0].Name,
		SpotifyImage: track.Album.Images[0].URL,
	}

	// Create a new Track record in the db
	newTrack := &model.Track{
		ISRC:         isrc,
		Title:        trackDetails.Title,
		ArtistName:   trackDetails.ArtistName,
		SpotifyImage: trackDetails.SpotifyImage,
	}
	err = ts.DatabaseAccessor.SaveTrack(newTrack)
	if err != nil {
		return nil, fmt.Errorf("failed to save track in the database")
	}
	return trackDetails, nil
}

// @Summary Search tracks by artist name
// @Description Search tracks from the database or Spotify by artist name
// @ID search-track-by-artist
// @Produce json
// @Param artistName path string true "Name of the artist"
// @Success 200 {array} model.TrackDetails
// @Failure 404 {object} map[string]interface{} "No tracks found for the artist"
// @Router /track/artist/{artistName} [get]
func (ts *TrackService) SearchTracksByArtistName(artistName string) ([]model.TrackDetails, error) {
	tracks, err := ts.DatabaseAccessor.GetTracksByArtistName(artistName)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tracks from the database")
	}

	if len(tracks) == 0 {
		return nil, fmt.Errorf("no tracks found for the artist")
	}

	var trackDetailsList []model.TrackDetails
	for _, track := range tracks {
		trackDetailsList = append(trackDetailsList, model.TrackDetails{
			ISRC:         track.ISRC,
			Title:        track.Title,
			ArtistName:   track.ArtistName,
			SpotifyImage: track.SpotifyImage,
		})
	}
	return trackDetailsList, nil
}

// @Summary Create a new track
// @Description Create a new track record in the database
// @ID create-track
// @Accept json
// @Produce json
// @Param trackDetails body model.TrackDetails true "Track details to create"
// @Success 200 {object} model.TrackDetails "Existing track details"
// @Success 201 {object} model.TrackDetails
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 409 {object} map[string]interface{} "Track with ISRC code already exists"
// @Router /track [post]
func (ts *TrackService) GetOrCreateTrackDetails(isrc string) (*model.TrackDetails, error) {

	existingTrack, err := ts.DatabaseAccessor.GetTrackByISRC(isrc)

	if err == nil {
		trackDetails := &model.TrackDetails{
			ISRC:         existingTrack.ISRC,
			Title:        existingTrack.Title,
			ArtistName:   existingTrack.ArtistName,
			SpotifyImage: existingTrack.SpotifyImage,
		}
		return trackDetails, nil
	}

	track, err := ts.SpotifyService.SearchTrackByISRC(isrc)
	if err != nil {
		return nil, fmt.Errorf("track not found")
	}

	trackDetails := &model.TrackDetails{
		ISRC:         isrc,
		Title:        track.Name,
		ArtistName:   track.Artists[0].Name,
		SpotifyImage: track.Album.Images[0].URL,
	}
	newTrack := &model.Track{
		ISRC:         isrc,
		Title:        trackDetails.Title,
		ArtistName:   trackDetails.ArtistName,
		SpotifyImage: trackDetails.SpotifyImage,
	}
	err = ts.DatabaseAccessor.SaveTrack(newTrack)
	if err != nil {
		return nil, fmt.Errorf("failed to save track in the database")
	}
	return trackDetails, nil
}

// @Summary Update a track by ISRC
// @Description Update an existing track record in the database by ISRC
// @Accept json
// @Produce json
// @Param isrc path string true "ISRC code of the track to be updated"
// @Param trackDetails body model.TrackDetails true "Updated track details"
// @Success 200 {object} model.TrackDetails
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /track/{isrc} [put]
func (ts *TrackService) UpdateTrack(isrc string, updatedTrackDetails *model.TrackDetails) (*model.TrackDetails, error) {
	existingTrack, err := ts.DatabaseAccessor.GetTrackByISRC(isrc)
	if err != nil {
		return nil, fmt.Errorf("track not found")
	}

	// Updating existing track fields with new values
	existingTrack.Title = updatedTrackDetails.Title
	existingTrack.ArtistName = updatedTrackDetails.ArtistName
	existingTrack.SpotifyImage = updatedTrackDetails.SpotifyImage

	// Saving updated details to db
	if err := ts.DatabaseAccessor.DB.Save(existingTrack).Error; err != nil {
		return nil, fmt.Errorf("failed to update track in the database")
	}

	// Return the updated track details
	return &model.TrackDetails{
		ISRC:         existingTrack.ISRC,
		Title:        existingTrack.Title,
		ArtistName:   existingTrack.ArtistName,
		SpotifyImage: existingTrack.SpotifyImage,
	}, nil
}
