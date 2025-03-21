package services

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/models"
	"gorm.io/gorm"
)

func GetSongs(c *gin.Context, db *gorm.DB) {
	const funcName = "GetSongs"

	// Проверка всех условий
	query := db.Model(&models.SongTable{})
	if group := c.Query("group"); group != "" {
		query = query.Where("\"group\" LIKE ?", "%"+group+"%")
		slog.Debug("query by group",
			slog.String("function_name", funcName))
	}
	if song := c.Query("song"); song != "" {
		query = query.Where("song LIKE ?", "%"+song+"%")
		slog.Debug("query by song",
			slog.String("function_name", funcName))
	}
	if releaseDate := c.Query("release_date"); releaseDate != "" {
		query = query.Where("release_date = ?", releaseDate)
		slog.Debug("query by release_date",
			slog.String("function_name", funcName))
	}
	if text := c.Query("text"); text != "" {
		query = query.Where("text LIKE ?", "%"+text+"%")
		slog.Debug("query by text",
			slog.String("function_name", funcName))
	}
	if link := c.Query("link"); link != "" {
		query = query.Where("link = ?", link)
		slog.Debug("query by link",
			slog.String("function_name", funcName))
	}

	// Получение страницы
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		slog.Error("invalid page format",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid page format"})
		return
	}

	// Получение лимита
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		slog.Error("invalid limit format",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid limit format"})
		return
	}
	offset := (page - 1) * limit
	slog.Debug("offset set",
		slog.String("function_name", funcName),
		slog.Int("offset", offset))

	// Получение общего количества
	var total int64
	query.Count(&total)
	var songs []models.SongResponse
	if err := query.Limit(limit).Offset(offset).Find(&songs).Error; err != nil {
		slog.Error("error song returned",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{Error: "Error fetching songs"})
		return
	}

	response := models.GetSongsResponse{
		Data:       songs,
		Page:       page,
		Limit:      limit,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		TotalItems: total,
	}
	c.JSON(http.StatusOK, response)
}

func GetSongLyrics(c *gin.Context, db *gorm.DB) {
	const funcName = "GetSongLyrics"

	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("invalid id",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	// Проверка существования
	var song models.SongTable
	if err := db.First(&song, songID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound,
				models.ErrorResponse{Error: "Song not found"})
		} else {
			slog.Error("failed to retrieve song",
				slog.String("function_name", funcName),
				slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError,
				models.ErrorResponse{Error: "Failed to retrieve song"})
		}
		return
	}

	if song.Text == "" {
		c.JSON(http.StatusOK,
			models.GetSongLyricsResponse{Verses: []string{},
				Page:        1,
				TotalVerses: 0})
		return
	}

	verses := strings.Split(song.Text, "\n\n")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		slog.Error("invalid page",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid page format"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		slog.Error("invalid limit",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid limit format"})
		return
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(verses) {
		c.JSON(http.StatusOK,
			models.GetSongLyricsResponse{Verses: []string{},
				Page:        page,
				Limit:       limit,
				TotalVerses: len(verses)})
		return
	}
	if end > len(verses) {
		end = len(verses)
	}

	response := models.GetSongLyricsResponse{
		Verses:      verses[start:end],
		Page:        page,
		Limit:       limit,
		TotalVerses: len(verses),
	}
	c.JSON(http.StatusOK, response)
}

func AddSong(c *gin.Context, cfg *config.Config, db *gorm.DB) {
	const funcName = "AddSong"

	var input struct {
		Group string `json:"group" binding:"required"`
		Song  string `json:"song" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		slog.Error("invalid input",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid data format"})
		return
	}

	// Получение ответа от внешнего сервера
	songDetails, err := fetchSongDetails(input.Group, input.Song, cfg)
	if err != nil {
		slog.Error("failed to fetch song details",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{Error: "Invalid fetched song data"})
		return
	}
	if songDetails == nil {
		slog.Error("fetched song details are nil",
			slog.String("function_name", funcName))
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{Error: "Song create error"})
		return
	}

	songToInsert := models.SongTable{
		Group:       songDetails.Group,
		Song:        songDetails.Song,
		ReleaseDate: songDetails.ReleaseDate,
		Text:        songDetails.Text,
		Link:        songDetails.Link,
	}
	if err := db.Create(&songToInsert).Error; err != nil {
		slog.Error("failed to create song",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{Error: "Song create error"})
		return
	}

	// Замена на ID в БД
	(*songDetails).ID = songToInsert.ID
	c.JSON(http.StatusCreated, songDetails)
	slog.Debug("new song added",
		slog.String("function_name", funcName),
		slog.String("data", fmt.Sprintf("%#v", songDetails)))
}

func UpdateSong(c *gin.Context, db *gorm.DB) {
	const funcName string = "UpdateSong"

	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("invalid song ID",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid song ID"})
		return
	}

	var song models.SongTable
	if err := db.First(&song, songID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound,
				models.ErrorResponse{Error: "Song not found"})
		} else {
			slog.Error("failed to retrieve song",
				slog.String("function_name", funcName),
				slog.String("error", err.Error()))
			c.JSON(http.StatusInternalServerError,
				models.ErrorResponse{Error: "Failed to retrieve song"})
		}
		return
	}

	var updateData models.UpdateSongRequest
	if err := c.ShouldBindJSON(&updateData); err != nil {
		slog.Error("invalid request body",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if updateData.Group != nil {
		song.Group = *updateData.Group
	}
	if updateData.Song != nil {
		song.Song = *updateData.Song
	}
	if updateData.ReleaseDate != nil {
		song.ReleaseDate = *updateData.ReleaseDate
	}
	if updateData.Text != nil {
		song.Text = *updateData.Text
	}
	if updateData.Link != nil {
		song.Link = *updateData.Link
	}
	if err := db.Save(&song).Error; err != nil {
		slog.Error("failed to update song",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{Error: "Failed to update song"})
		return
	}

	answerSong := models.SongResponse{
		ID:          song.ID,
		Group:       song.Group,
		Song:        song.Song,
		ReleaseDate: song.ReleaseDate,
		Text:        song.Text,
		Link:        song.Link,
	}
	c.JSON(http.StatusOK, answerSong)
	slog.Debug("Song updated",
		slog.String("function_name", funcName),
		slog.Int("id", songID))
}

func DeleteSong(c *gin.Context, db *gorm.DB) {
	const funcName = "DeleteSong"

	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		slog.Error("invalid song format",
			slog.String("function_name", funcName),
			slog.Int("id", songID),
			slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var song models.SongTable
	if err := db.First(&song, songID).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			slog.Error("song not found",
				slog.String("function_name", funcName),
				slog.String("error", err.Error()),
				slog.Int("song_id", songID),
			)
			c.JSON(http.StatusNotFound,
				models.ErrorResponse{Error: "Song not found"})
		default:
			slog.Error("failed to retrieve song",
				slog.String("function_name", funcName),
				slog.String("error", err.Error()),
				slog.Int("song_id", songID),
			)
			c.JSON(http.StatusInternalServerError,
				models.ErrorResponse{Error: "Failed to retrieve song"})
		}
		return
	}

	if err := db.Delete(&models.SongTable{}, songID).Error; err != nil {
		slog.Error("failed to delete song",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{Error: "Failed to delete song"})
		return
	}

	c.JSON(http.StatusOK,
		models.DeleteSongResponse{Message: "Song deleted"})
	slog.Debug("song deleted",
		slog.String("function_name", funcName),
		slog.Int("id", songID))
}
