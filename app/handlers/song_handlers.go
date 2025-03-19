package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/database"
	"github.com/mikhail-kirillov/Music-info/services"
)

// @Summary		Получить список песен
// @Description	Возвращает список песен с фильтрацией и пагинацией.
// @Tags			songs
// @Produce		json
// @Param			group			query		string	false	"Название группы"
// @Param			song			query		string	false	"Название песни"
// @Param			release_date	query		string	false	"Дата релиза"
// @Param			text			query		string	false	"Текст песни"
// @Param			link			query		string	false	"Ссылка на песню"
// @Param			page			query		int		false	"Номер страницы"					default(1)
// @Param			limit			query		int		false	"Количество элементов на странице"	default(10)
// @Success		200				{object}	models.GetSongsResponse
// @Failure		400				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router			/songs [get]
func GetSongs(c *gin.Context, db database.Database) {
	services.GetSongs(c, db)
}

// @Summary		Получить текст песни
// @Description	Возвращает текст песни построчно с пагинацией.
// @Tags			songs
// @Produce		json
// @Param			id		path		int	true	"ID песни"
// @Param			page	query		int	false	"Номер страницы"				default(1)
// @Param			limit	query		int	false	"Количество строк на странице"	default(10)
// @Success		200		{object}	models.GetSongLyricsResponse
// @Failure		400		{object}	models.ErrorResponse
// @Failure		404		{object}	models.ErrorResponse
// @Router			/songs/{id}/lyrics [get]
func GetSongLyrics(c *gin.Context, db database.Database) {
	services.GetSongLyrics(c, db)
}

// @Summary		Добавить песню
// @Description	Добавляет песню на основе предоставленных данных о группе и названии песни.
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			input	body		models.AddSongRequest	true	"Данные новой песни"
// @Success		201		{object}	models.Song
// @Failure		400		{object}	models.ErrorResponse
// @Failure		500		{object}	models.ErrorResponse
// @Router			/songs [post]
func AddSong(c *gin.Context, cfg *config.Config, db database.Database) {
	services.AddSong(c, cfg, db)
}

// @Summary		Обновить песню
// @Description	Обновляет название, группу, дату релиза, текст или ссылку на песню.
// @Tags			songs
// @Accept			json
// @Produce		json
// @Param			id		path		int							true	"ID песни"
// @Param			input	body		models.UpdateSongRequest	true	"Обновленные данные песни"
// @Success		200		{object}	models.Song
// @Failure		400		{object}	models.ErrorResponse
// @Failure		404		{object}	models.ErrorResponse
// @Failure		500		{object}	models.ErrorResponse
// @Router			/songs/{id} [put]
func UpdateSong(c *gin.Context, db database.Database) {
	services.UpdateSong(c, db)
}

// @Summary		Удалить песню
// @Description	Удаляет песню из базы данных по ее идентификатору.
// @Tags			songs
// @Param			id	path		int	true	"ID песни"
// @Success		200	{object}	models.DeleteSongResponse
// @Failure		400	{object}	models.ErrorResponse
// @Failure		404	{object}	models.ErrorResponse
// @Failure		500	{object}	models.ErrorResponse
// @Router			/songs/{id} [delete]
func DeleteSong(c *gin.Context, db database.Database) {
	services.DeleteSong(c, db)
}
