package data

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID         string    `gorm:"type:varchar(128)"`
	Caption    string    `gorm:"type:text"`
	Location   string    `gorm:"type:text"`
	Width      int       `gorm:"type:integer;not null"`
	Height     int       `gorm:"type:integer;not null"`
	CreatedAt  time.Time `gorm:"not null;default:now();index"`
	Camera     string    `gorm:"type:text"`
	Lens       string    `gorm:"type:text"`
	Film       string    `gorm:"type:text"`
	ImageTypes []ImageType
	Tags       []Tag `gorm:"many2many:image_tags;"`
}

type ImageType struct {
	ImageID string        `gorm:"primarykey;not null"`
	Type    ImageTypeName `gorm:"primarykey;type:varchar(128);not null"`
}

type ImageTypeName string

const (
	PhotoADayImageType ImageTypeName = "photo-a-day"
)

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

func (s *ds) GetYearImages(year int, before time.Time, limit int) ([]*Image, error) {
	var images []*Image
	tx := s.db.
		Joins("JOIN image_types it ON it.image_id = images.id").
		Where("it.type = ?", PhotoADayImageType).
		Where("created_at < ?", before).
		Where("date_part('year', created_at) = ?", year).
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
	tags := image.Tags
	image.Tags = nil

	if tx := s.db.Create(image); tx.Error != nil {
		return tx.Error
	}

	for i := range tags {
		tag := &tags[i]
		var existing Tag
		result := s.db.Where("slug = ?", tag.Slug).First(&existing)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := s.db.Create(tag).Error; err != nil {
					return err
				}
			} else {
				return result.Error
			}
		} else {
			tags[i] = existing
		}
	}

	if len(tags) > 0 {
		return s.db.Model(image).Association("Tags").Append(tags)
	}
	return nil
}
