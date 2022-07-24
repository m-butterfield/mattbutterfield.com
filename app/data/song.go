package data

import (
	"time"
)

type Song struct {
	ID          string     `gorm:"type:varchar(255)"`
	Description string     `gorm:"type:text"`
	CreatedAt   *time.Time `gorm:"not null;default:now()"`
}

func (s *ds) GetSongs() ([]*Song, error) {
	var songs []*Song
	tx := s.db.
		Order("created_at DESC").
		Find(&songs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return songs, nil
}

func (s *ds) SaveSong(song *Song) error {
	if tx := s.db.Create(song); tx.Error != nil {
		return tx.Error
	}
	return nil
}
