package data

import (
	"strings"
	"time"
)

type Tag struct {
	Name   string  `gorm:"type:varchar(128);not null"`
	Slug   string  `gorm:"primarykey;type:varchar(128);not null"`
	Images []Image `gorm:"many2many:image_tags;"`
}

func MakeTagSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(name), " ", "-"))
}

func (s *ds) GetImagesByTag(slugs []string, before time.Time, limit int) ([]*Image, error) {
	var images []*Image
	tx := s.db.
		Joins("JOIN image_tags ON image_tags.image_id = images.id").
		Where("image_tags.tag_slug IN ?", slugs).
		Where("created_at < ?", before).
		Group("images.id").
		Having("COUNT(DISTINCT image_tags.tag_slug) = ?", len(slugs)).
		Order("created_at DESC").
		Limit(limit).
		Find(&images)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return images, nil
}

func (s *ds) GetTagsBySlugs(slugs []string) ([]*Tag, error) {
	var tags []*Tag
	tx := s.db.Where("slug IN ?", slugs).Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
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
