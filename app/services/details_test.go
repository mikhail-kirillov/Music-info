package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/models"
	"github.com/stretchr/testify/assert"
)

func TestFetchSongDetails_Success(t *testing.T) {
	expectedSong := &models.Song{
		ID:    1,
		Song:  "Test Song",
		Group: "Test Group",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/info", r.URL.Path)
		assert.Equal(t, "Test Group", r.URL.Query().Get("group"))
		assert.Equal(t, "Test Song", r.URL.Query().Get("song"))

		var testSong models.Song = models.Song{
			ID:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   time.Now(),
			Group:       "Test Group",
			Song:        "Test Song",
			ReleaseDate: "Test Date",
			Text:        "Test Text",
			Link:        "Test Link",
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(testSong)
	}))
	defer server.Close()

	cfg := &config.Config{MusicApiURL: server.URL}

	song, err := fetchSongDetails("Test Group", "Test Song", cfg)

	if err != nil {
		t.Fatalf("Error call fetchSongDetails: %v", err)
	}

	assert.NotNil(t, song, "Must back object Song")
	assert.Equal(t, expectedSong.ID, song.ID)
	assert.Equal(t, expectedSong.Song, song.Song)
	assert.Equal(t, expectedSong.Group, song.Group)
}

func TestFetchSongDetails_RequestError(t *testing.T) {
	cfg := &config.Config{MusicApiURL: "http://invalid-url"}

	song, err := fetchSongDetails("Group", "Song", cfg)

	assert.Error(t, err, "Must be request error")
	assert.Nil(t, song, "Song must be nil")
}

func TestFetchSongDetails_StatusError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	cfg := &config.Config{MusicApiURL: server.URL}

	song, err := fetchSongDetails("Group", "Song", cfg)

	assert.Error(t, err, "Must be request error")
	assert.Nil(t, song, "Song must be nil")
}

func TestFetchSongDetails_JSONDecodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	cfg := &config.Config{MusicApiURL: server.URL}

	song, err := fetchSongDetails("Group", "Song", cfg)

	assert.Error(t, err, "Must be JSON decode error")
	assert.Nil(t, song, "Song must be nil")
}
