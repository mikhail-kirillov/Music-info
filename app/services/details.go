package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/models"
)

func fetchSongDetails(group, song string, cfg *config.Config) (*models.SongResponse, error) {
	const funcName string = "fetchSongDetails"

	url := fmt.Sprintf("%s/info?group=%s&song=%s",
		cfg.MusicApiURL,
		url.QueryEscape(group),
		url.QueryEscape(song))
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("get request error",
			slog.String("function_name", funcName),
			slog.String("url", url),
			slog.String("error", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()
	slog.Debug("request was received", slog.String("function_name", funcName))

	if resp.StatusCode != http.StatusOK {
		slog.Error("get request status error",
			slog.String("function_name", funcName),
			slog.Int("status_code", resp.StatusCode))
		return nil, errors.New("request to API")
	}

	var songDetail models.SongResponse
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		slog.Error("decode error",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		return nil, err
	}
	songDetail.Group = group
	songDetail.Song = song
	slog.Debug("data returned",
		slog.String("function_name", funcName),
		slog.String("body", fmt.Sprintf("%#v", songDetail)))

	return &songDetail, nil
}
