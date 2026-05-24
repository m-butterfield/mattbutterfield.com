package data

func Migrate() error {
	s, err := getDS()
	if err != nil {
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
