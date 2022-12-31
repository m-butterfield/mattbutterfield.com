package data

func Migrate() error {
	s, err := getDS()
	if err != nil {
		return err
	}
	err = s.db.AutoMigrate(
		&Image{},
		&ImageType{},
		&Song{},
	)
	if err != nil {
		return err
	}
	return nil
}
