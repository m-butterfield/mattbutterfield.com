package data

import (
	"time"
)

type StravaAccessToken struct {
	ID           string
	AccessToken  string    `gorm:"not null;type:varchar(40)"`
	RefreshToken string    `gorm:"not null;type:varchar(40)"`
	Expiry       time.Time `gorm:"not null"`
}

func (s *ds) GetStravaAccessToken(id string) (*StravaAccessToken, error) {
	token := &StravaAccessToken{}
	tx := s.db.First(token, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return token, nil
}

func (s *ds) CreateStravaAccessToken(token *StravaAccessToken) error {
	if tx := s.db.Create(token); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) UpdateStravaAccessToken(token *StravaAccessToken) error {
	if tx := s.db.Save(token); tx.Error != nil {
		return tx.Error
	}
	return nil
}
