package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RetroPage(c echo.Context) error {

	retro_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	if _, err := utils.ReadCookie(c, "name"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/guest?next=/retro/"+retro_id.String())
	}
	retro, err := models.FetchRetro(retro_id)
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	template, err := models.FetchTemplate(retro.Template)
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	records, _ := models.FetchRecordsByRetro(retro_id)
	categories := strings.Split(template.Categories, ", ")
	context := make(map[string][]models.Record)
	c_ids := make(map[string]string)

	for _, category := range categories {
		context[category] = []models.Record{}

		regex := regexp.MustCompile(`[^a-zA-Z_\-]`)
		c_ids[category] = regex.ReplaceAllString(category, "_")
		if retro.Visible {
			for _, record := range records {
				if record.Category == category {
					context[category] = append(context[category], record)
				}
			}
		}
	}
	return views.RetroPage(context, c_ids, "Retro from "+string(retro.Created)).Render(c.Request().Context(), c.Response().Writer)
}

func RetroMakeVisible(c echo.Context) error {
	if _, err := utils.ReadCookie(c, "user"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	retro_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	visible := c.FormValue("visible")
	_ = !(visible == "true")

	return c.Redirect(http.StatusSeeOther, "/retro/"+retro_id.String())
}

func RetroItemCreate(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")
	id := c.Param("id")
	retro_id := uuid.MustParse(id)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/guest?next=/retro/"+retro_id.String())
	}
	category := c.FormValue("category")
	content := c.FormValue("content")
	var record models.Record
	record.Id = uuid.New()
	record.Retro = retro_id
	record.Author = name.Value
	record.Category = category
	record.Content = content
	if err = models.CreateRecord(&record); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return views.RetroItem(record.Content, record.Author).Render(c.Request().Context(), c.Response().Writer)
}

func RetroCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	c.Request().ParseForm()
	log.Error(c.Request().Form)
	template_id := c.FormValue("template_id")
	new_retro := models.Retro{}
	new_retro.Id = uuid.New()
	new_retro.User = uuid.MustParse(user.Value)
	new_retro.Template = uuid.MustParse(template_id)
	new_retro.Created = time.Now().Format("2006-01-02")
	new_retro.Visible = false
	if err := models.CreateRetro(&new_retro); err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return c.Redirect(http.StatusSeeOther, "/retro/"+new_retro.Id.String())
}
