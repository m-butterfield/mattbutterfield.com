package data

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestGetStravaAccessToken(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	token := &StravaAccessToken{
		ID:           "main",
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		Expiry:       time.Now(),
	}
	if err = s.CreateStravaAccessToken(token); err != nil {
		t.Fatal(err)
	}

	var result *StravaAccessToken
	if result, err = s.GetStravaAccessToken("main"); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, token.ID, result.ID)
	assert.Equal(t, token.AccessToken, result.AccessToken)
	assert.Equal(t, token.RefreshToken, result.RefreshToken)
	assert.Equal(t, token.Expiry.Unix(), result.Expiry.Unix())

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaAccessToken{})
}

func TestUpdateStravaAccessToken(t *testing.T) {
	s, err := getDS()
	if err != nil {
		t.Fatal(err)
	}
	token := &StravaAccessToken{
		ID:           "main",
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		Expiry:       time.Now(),
	}
	if err = s.CreateStravaAccessToken(token); err != nil {
		t.Fatal(err)
	}

	token.AccessToken = "new_access_token"
	token.RefreshToken = "new_refresh_token"
	token.Expiry = time.Now().Add(time.Hour)

	if err = s.UpdateStravaAccessToken(token); err != nil {
		t.Fatal(err)
	}

	var result *StravaAccessToken
	if result, err = s.GetStravaAccessToken("main"); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, token.ID, result.ID)
	assert.Equal(t, token.AccessToken, result.AccessToken)
	assert.Equal(t, token.RefreshToken, result.RefreshToken)
	assert.Equal(t, token.Expiry.Unix(), result.Expiry.Unix())

	s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&StravaAccessToken{})
}
