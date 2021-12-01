package data

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

const (
	baseSelectSongRegex = "^SELECT id, description, created_at FROM songs "
	getSongsRegex       = baseSelectSongRegex + " ORDER BY created_at DESC"
	saveSongsRegex      = "^INSERT INTO songs VALUES \\(\\$1, \\$2, CURRENT_TIMESTAMP\\)"
)

func TestGetSongs(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}

	id, description, createdAt := "20901202", "drone", time.Date(2021, time.Month(9), 6, 13, 11, 0, 0, time.UTC)
	dbMock.ExpectQuery(getSongsRegex).
		WillReturnRows(sqlmock.NewRows([]string{"id", "description", "created_at"}).AddRow(id, description, &createdAt))

	songs, err := store.GetSongs()
	if err != nil {
		t.Fatal(err)
	}
	err = dbMock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled database expectations: %s", err)
	}
	if len(songs) != 1 {
		t.Fatalf("Unexpected length of songs: %d", len(songs))
	}
	song := songs[0]
	if song.ID != id {
		t.Errorf("Unexpected song id: %s != %s", id, song.ID)
	}
	if song.Description != description {
		t.Errorf("Unexpected song description: %s != %s", description, song.Description)
	}
	if song.CreatedAt != &createdAt {
		t.Errorf("Unexpected song createdAt: %s != %s", createdAt, song.CreatedAt)
	}
}

func TestSaveSong(t *testing.T) {
	db, dbMock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	store := &dbStore{db: db}
	dbMock.ExpectPrepare(saveSongsRegex).ExpectExec().WithArgs("testID", "testDescription").
		WillReturnResult(sqlmock.NewResult(1, 1))
	err = store.SaveSong("testID", "testDescription")
	if err != nil {
		t.Fatal(err)
	}
}
