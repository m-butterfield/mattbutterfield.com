package data

func Migrate() error {
	s, err := getDS()
	if err != nil {
		return err
	}
	if err := s.db.Migrator().DropTable(&Tag{}); err != nil {
		return err
	}
	if err := s.db.Exec("DROP TABLE IF EXISTS image_tags").Error; err != nil {
		return err
	}
	err = s.db.AutoMigrate(
		&Image{},
		&ImageType{},
		&Tag{},
		&StravaAccessToken{},
		&StravaActivity{},
		&Song{},
		&MapBoxConfig{},
	)
	if err != nil {
		return err
	}
	return nil
}
