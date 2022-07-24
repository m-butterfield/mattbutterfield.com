package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignedUploadURL(t *testing.T) {
	body, err := json.Marshal(&signedUploadURLRequest{
		FileName:    "test.wav?123456",
		ContentType: "audio/wav",
	})
	r, err := http.NewRequest(http.MethodPost, "/admin/signed_upload_url", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.AddCookie(&http.Cookie{Name: "auth", Value: "1234"})
	authArray = []byte("1234")

	w := httptest.NewRecorder()
	testRouter().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Unexpected return code: %d", w.Code)
	}
	result := &signedUploadURLResponse{}
	if err := json.NewDecoder(w.Body).Decode(result); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(result.URL, "https://storage.googleapis.com/files.mattbutterfield.com/uploads/test.wav") {
		t.Errorf("Unexpected URL result: %s", result.URL)
	}
}
