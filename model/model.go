package model

type Track struct {
	ID           uint `gorm:"primary_key"`
	ISRC         string
	Title        string
	ArtistName   string
	SpotifyImage string
}

type TrackDetails struct {
	ISRC         string `json:"isrc"`
	Title        string `json:"title"`
	ArtistName   string `json:"artist_name"`
	SpotifyImage string `json:"spotify_image"`
}
