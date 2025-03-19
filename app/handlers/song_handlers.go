package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/services"
	"gorm.io/gorm"
)

// @Summary		Get song list
// @Description	Returns a list of songs with filtering and pagination
// @Tags			songs
// @Produce		json
// @Param			group			query		string	false	"Group name"
// @Param			song			query		string	false	"Song Title"
// @Param			release_date	query		string	false	"Release date"
// @Param			text			query		string	false	"Song lyrics"
// @Param			link			query		string	false	"Link to the song"
// @Param			page			query		int		false	"Page number"					default(1)
// @Param			limit			query		int		false	"Number of elements per page"	default(10)
// @Success		200				{object}	models.GetSongsResponse
// @Failure		400				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router			/songs [get]
func GetSongs(c *gin.Context, db *gorm.DB) {
	services.GetSongs(c, db)
}

// @Summary		Get song lyrics
// @Description	Returns the lyrics line by line with pagination
// @Tags			songs
// @Produce		json
// @Param			id		path		int	true	"Song ID"
// @Param			page	query		int	false	"Verse number"	default(1)
// @Param			limit	query		int	false	"Limit number"	default(10)
// @Success		200		{object}	models.GetSongLyricsResponse
// @Failure		400		{object}	models.ErrorResponse
// @Failure		404		{object}	models.ErrorResponse
// @Router			/songs/{id}/lyrics [get]
func GetSongLyrics(c *gin.Context, db *gorm.DB) {
	services.GetSongLyrics(c, db)
}

// @Summary		Add song
// @Description	Adds a song based on the provided band and song title information
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			input	body		models.AddSongRequest	true	"New song details"
// @Success		201		{object}	models.SongResponse
// @Failure		400		{object}	models.ErrorResponse
// @Failure		500		{object}	models.ErrorResponse
// @Router			/songs [post]
func AddSong(c *gin.Context, cfg *config.Config, db *gorm.DB) {
	services.AddSong(c, cfg, db)
}

// @Summary		Update song
// @Description	Updates the title, band, release date, lyrics or link to the song
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			id		path		int							true	"Song ID"
// @Param			input	body		models.UpdateSongRequest	true	"Updated song data"
// @Success		200		{object}	models.SongResponse
// @Failure		400		{object}	models.ErrorResponse
// @Failure		404		{object}	models.ErrorResponse
// @Failure		500		{object}	models.ErrorResponse
// @Router			/songs/{id} [put]
func UpdateSong(c *gin.Context, db *gorm.DB) {
	services.UpdateSong(c, db)
}

// @Summary		Delete song
// @Description	Removes a song from the database by its ID
// @Tags			songs
// @Param			id	path		int	true	"Song ID"
// @Success		200	{object}	models.DeleteSongResponse
// @Failure		400	{object}	models.ErrorResponse
// @Failure		404	{object}	models.ErrorResponse
// @Failure		500	{object}	models.ErrorResponse
// @Router			/songs/{id} [delete]
func DeleteSong(c *gin.Context, db *gorm.DB) {
	services.DeleteSong(c, db)
}
