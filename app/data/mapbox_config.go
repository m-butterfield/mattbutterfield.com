package data

type MapBoxConfig struct {
	Name               string `gorm:"primarykey"`
	HeatMapTilesetName string `gorm:"not null"`
}

const mapBoxConfigName = "main"

func (s *ds) GetMapBoxConfig() (*MapBoxConfig, error) {
	config := &MapBoxConfig{}
	tx := s.db.First(config, "name = ?", mapBoxConfigName)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return config, nil
}

func (s *ds) UpdateMapBoxConfig(config *MapBoxConfig) error {
	if tx := s.db.Model(&config).Where("name = ?", mapBoxConfigName).Updates(&config); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *ds) CreateMapBoxConfig(config *MapBoxConfig) error {
	if tx := s.db.Create(config); tx.Error != nil {
		return tx.Error
	}
	return nil
}
