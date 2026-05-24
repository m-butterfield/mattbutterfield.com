package data

import (
	"time"
)

type Tag struct {
	Name   string  `gorm:"primarykey;type:varchar(128);not null"`
	Images []Image `gorm:"many2many:image_tags;"`
}

func (s *ds) GetImagesByTag(names []string, before time.Time, limit int) ([]*Image, error) {
	var images []*Image
	tx := s.db.
		Joins("JOIN image_tags ON image_tags.image_id = images.id").
		Where("image_tags.tag_name IN ?", names).
		Where("created_at < ?", before).
		Group("images.id").
		Having("COUNT(DISTINCT image_tags.tag_name) = ?", len(names)).
		Order("created_at DESC").
		Limit(limit).
		Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return images, nil
}

func (s *ds) GetAllTags() ([]*Tag, error) {
	var tags []*Tag
	tx := s.db.Order("name ASC").Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *ds) GetTagsByNames(names []string) ([]*Tag, error) {
	var tags []*Tag
	tx := s.db.Where("name IN ?", names).Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *ds) GetImageTags(imageID string) ([]*Tag, error) {
	var tags []*Tag
	tx := s.db.
		Joins("JOIN image_tags ON image_tags.tag_name = tags.name").
		Where("image_tags.image_id = ?", imageID).
		Order("tags.name ASC").
		Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *ds) AddImageTag(imageID string, tagName string) error {
	tag := &Tag{Name: tagName}
	tx := s.db.Where("name = ?", tagName).FirstOrCreate(tag)
	if tx.Error != nil {
		return tx.Error
	}
	return s.db.Model(&Image{ID: imageID}).Association("Tags").Append(tag)
}

func (s *ds) RemoveImageTag(imageID string, tagName string) error {
	tx := s.db.Model(&Image{ID: imageID}).Association("Tags").Delete(&Tag{Name: tagName})
	return tx
}
