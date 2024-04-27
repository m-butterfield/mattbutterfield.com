package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestGetMapboxConfig(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	config := &MapBoxConfig{
		Name:               "main",
		HeatMapTilesetName: "test tileset",
	}
	if err = s.CreateMapBoxConfig(config); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetMapBoxConfig()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config.HeatMapTilesetName, result.HeatMapTilesetName)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(MapBoxConfig{})
}

func TestUpdateMapboxConfig(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	config := &MapBoxConfig{
		Name:               "main",
		HeatMapTilesetName: "test tileset",
	}
	if err = s.CreateMapBoxConfig(config); err != nil {
		t.Fatal(err)
	}

	config.HeatMapTilesetName = "new tileset"
	if err = s.UpdateMapBoxConfig(config); err != nil {
		t.Fatal(err)
	}

	result, err := s.GetMapBoxConfig()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config.HeatMapTilesetName, result.HeatMapTilesetName)

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(MapBoxConfig{})
}
