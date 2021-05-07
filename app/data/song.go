package data

import (
	"database/sql"
	"log"
	"time"
)

const (
	baseSelectSongQuery = "SELECT id, description, created_at FROM songs "
	getSongsQuery       = baseSelectSongQuery + " ORDER BY created_at DESC"
)

type Song struct {
	ID          string
	Description string
	CreatedAt   *time.Time
}

func (s *dbStore) GetSongs() ([]*Song, error) {
	rows, err := s.db.Query(getSongsQuery)
	if err != nil {
		return nil, err
	}
	var songs []*Song
	for rows.Next() {
		var (
			description sql.NullString
			song        = &Song{}
		)
		err := rows.Scan(&song.ID, &description, &song.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		song.Description = description.String
		songs = append(songs, song)
	}
	return songs, nil
}
