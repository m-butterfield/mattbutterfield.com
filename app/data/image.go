package data

import (
	"time"
)

type Image struct {
	ID        string    `gorm:"type:varchar(128)"`
	Caption   string    `gorm:"type:text"`
	Location  string    `gorm:"type:text"`
	Width     int       `gorm:"type:integer;not null"`
	Height    int       `gorm:"type:integer;not null"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (s *ds) GetImage(id string) (*Image, error) {
	image := &Image{}
	tx := s.db.First(&image, "id = $1", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return image, nil
}

func (s *ds) GetImages(before time.Time, limit int) ([]*Image, error) {
	var images []*Image
	tx := s.db.
		Where("created_at < $1", before).
		Order("created_at DESC").
		Limit(limit).
		Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return images, nil
}

func (s *ds) GetRandomImage() (*Image, error) {
	image := &Image{}
	tx := s.db.First(&image, "id = (SELECT id FROM images ORDER BY RANDOM() LIMIT 1)")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return image, nil
}

func (s *ds) SaveImage(image *Image) error {
	if tx := s.db.Create(image); tx.Error != nil {
		return tx.Error
	}
	return nil
}
