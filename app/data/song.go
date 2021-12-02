package data

import (
	"database/sql"
	"log"
	"time"
)

const (
	baseSelectSongQuery = "SELECT id, description, created_at FROM songs "
	getSongsQuery       = baseSelectSongQuery + " ORDER BY created_at DESC"
	insertSongQuery     = "INSERT INTO songs VALUES ($1, $2, CURRENT_TIMESTAMP)"
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

func (s *dbStore) SaveSong(id, description string) error {
	stmt, err := s.db.Prepare(insertSongQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, nullString(description))
	if err != nil {
		return err
	}
	return nil
}

func nullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}