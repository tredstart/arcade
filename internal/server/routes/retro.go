package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RetroPage(c echo.Context) error {

	retro_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
    if _, err := utils.ReadCookie(c, "name"); err != nil {
        return c.Redirect(http.StatusSeeOther, "/guest?next=/retro/" + retro_id.String())
    }
	retro, err := models.FetchRetro(retro_id)
	if err != nil {
		return views.ErrorBlock("Couldn't find this retro").Render(c.Request().Context(), c.Response().Writer)
	}
	template, err := models.FetchTemplate(retro.Template)
	if err != nil {
		return views.ErrorBlock("Couldn't find this template").Render(c.Request().Context(), c.Response().Writer)
	}

	records, _ := models.FetchRecordsByRetro(retro_id)
	categories := strings.Split(template.Categories, ", ")
	context := make(map[string][]models.Record)
	c_ids := make(map[string]string)

	user, _ := utils.ReadCookie(c, "user")
	for _, category := range categories {
		context[category] = []models.Record{}

		regex := regexp.MustCompile(`[^a-zA-Z0-9_\-]`)
		c_ids[category] = regex.ReplaceAllString(category, "_")
		if user != nil {
			for _, record := range records {
				if record.Category == category {
					context[category] = append(context[category], record)
				}
			}
		}
	}
	return views.RetroPage(context, c_ids, "Retro from " + string(retro.Created)).Render(c.Request().Context(), c.Response().Writer)
}

func RetroItemCreate(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")
	if err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	id := c.Param("id")
	retro_id := uuid.MustParse(id)
	category := c.FormValue("category")
	content := c.FormValue("content")
	if len(content) == 0 {
		return views.ErrorBlock("Cannot create item with empty content").Render(c.Request().Context(), c.Response().Writer)
	}
	var record models.Record
	record.Id = uuid.New()
	record.Retro = retro_id
	record.Author = name.Value
	record.Category = category
	record.Content = content
	if err = models.CreateRecord(&record); err != nil {
		log.Println("Error while creating new record: ", err)
		return views.ErrorBlock("Couldn't create new record").Render(c.Request().Context(), c.Response().Writer)
	}
	return views.RetroItem(record.Content, record.Author).Render(c.Request().Context(), c.Response().Writer)
}

func RetroCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	template_id := c.FormValue("template_id")
	new_retro := models.Retro{}
	new_retro.Id = uuid.New()
	new_retro.User = uuid.MustParse(user.Value)
	new_retro.Template = uuid.MustParse(template_id)
	new_retro.Created = time.Now().Format("2006-01-02")
	if err := models.CreateRetro(&new_retro); err != nil {
		log.Println("Error while creating new retro: ", err)
		return views.ErrorBlock("Couldn't create new retro").Render(c.Request().Context(), c.Response().Writer)
	}
	return c.Redirect(http.StatusSeeOther, "/retro/" + new_retro.Id.String())
}


