package data

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	Name   string  `gorm:"type:varchar(128);not null"`
	Slug   string  `gorm:"primarykey;type:varchar(128);not null"`
	Images []Image `gorm:"many2many:image_tags;"`
}

func MakeTagSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(name), " ", "-"))
}

func (s *ds) GetTags() ([]*Tag, error) {
	var tags []*Tag
	tx := s.db.Order("name ASC").Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *ds) GetImagesByTag(slug string, before time.Time, limit int) ([]*Image, error) {
	var images []*Image
	tx := s.db.
		Joins("JOIN image_tags ON image_tags.image_id = images.id").
		Where("image_tags.tag_slug = ?", slug).
		Where("created_at < ?", before).
		Order("created_at DESC").
		Limit(limit).
		Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return images, nil
}

func (s *ds) GetImageTags(imageID string) ([]*Tag, error) {
	var tags []*Tag
	tx := s.db.
		Joins("JOIN image_tags ON image_tags.tag_slug = tags.slug").
		Where("image_tags.image_id = ?", imageID).
		Order("tags.name ASC").
		Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *ds) AddImageTag(imageID string, tagName string) error {
	slug := MakeTagSlug(tagName)
	tag := &Tag{Name: tagName, Slug: slug}
	tx := s.db.Where("slug = ?", slug).FirstOrCreate(tag)
	if tx.Error != nil {
		return tx.Error
	}
	return s.db.Model(&Image{ID: imageID}).Association("Tags").Append(tag)
}

func (s *ds) RemoveImageTag(imageID string, tagSlug string) error {
	tx := s.db.Model(&Image{ID: imageID}).Association("Tags").Delete(&Tag{Slug: tagSlug})
	return tx
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
