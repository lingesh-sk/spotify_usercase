package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/lingesh-sk/spotify_usercase/model"
)

type DatabaseAccessor struct {
	DB *gorm.DB
}

func NewDatabaseAccessor(db *gorm.DB) *DatabaseAccessor {
	return &DatabaseAccessor{DB: db}
}

func (d *DatabaseAccessor) SaveTrack(track *model.Track) error {
	return d.DB.Create(track).Error
}

// It retrieves a track from the db by ISRC code
func (d *DatabaseAccessor) GetTrackByISRC(isrc string) (*model.Track, error) {
	var track model.Track
	err := d.DB.Where("isrc = ?", isrc).First(&track).Error
	return &track, err
}

// It retrieves tracks from the db by artist name
func (d *DatabaseAccessor) GetTracksByArtistName(artistName string) ([]model.Track, error) {
	var tracks []model.Track
	err := d.DB.Where("artist_name = ?", artistName).Find(&tracks).Error
	return tracks, err
}
