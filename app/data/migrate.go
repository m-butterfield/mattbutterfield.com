package data

func Migrate() error {
	s, err := getDS()
	if err != nil {
		return err
	}
	err = s.db.AutoMigrate(
		&Image{},
		&Song{},
	)
	if err != nil {
		return err
	}
	return nil
}
