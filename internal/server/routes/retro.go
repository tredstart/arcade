package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RetroPage(c echo.Context) error {
	retro_id, err := uuid.Parse(c.QueryParam("id"))
	if err != nil {
		return err
	}
	retro, err := models.FetchRetro(retro_id)
	if err != nil {
		return err
	}
	template, err := models.FetchTemplate(retro.Template)
	if err != nil {
		return err
	}

	records, _ := models.FetchRecordsByRetro(retro_id)
	categories := strings.Split(template.Categories, ", ")
	context := make(map[string][]models.Record)

	user, _ := utils.ReadCookie(c, "user")
	for _, category := range categories {
		context[category] = []models.Record{}
		if user != nil {
			for _, record := range records {
				if record.Category == category {
					context[category] = append(context[category], record)
				}
			}
		}
	}
	return nil
}

func RetroItemCreate(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")
	if err != nil {
		return err
	}
	retro_id := uuid.MustParse(c.QueryParam("id"))
	category := c.FormValue("category")
	content := c.FormValue("content")
	var record models.Record
	record.Id = uuid.New()
	record.Retro = retro_id
	record.Author = name.Value
	record.Category = category
	record.Content = content
    if err = models.CreateRecord(&record); err != nil {
		log.Println("Error while creating new record: ", err)
		return c.String(http.StatusTeapot, "Couldn't create new record")
    }
	return c.String(http.StatusCreated, "Record created successfuly")
}

func RetroCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return err
	}
	template_id := c.FormValue("template_id")
	new_retro := models.Retro{}
	new_retro.Id = uuid.New()
	new_retro.User = uuid.MustParse(user.Value)
	new_retro.Template = uuid.MustParse(template_id)
	new_retro.Created = time.Now().Format("2006-01-02")
	if err := models.CreateRetro(&new_retro); err != nil {
		log.Println("Error while creating new retro: ", err)
		return c.String(http.StatusTeapot, "Couldn't create new retro")
	}
	return c.String(http.StatusCreated, "Retro created successfuly")
}

func Templates(c echo.Context) error {
	// render a list of templates for some user
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return err
	}
	templates, _ := models.FetchTemplatesByUser(user.Value)
	log.Println(templates)
	return views.Templates(templates).Render(c.Request().Context(), c.Response().Writer)
}

func TemplatesNew(c echo.Context) error {
	// render a form for new tempalates
	_, err := utils.ReadCookie(c, "user")
	if err != nil {
		return err
	}
	return views.CreateTemplateForm().Render(c.Request().Context(), c.Response().Writer)
}

func TempalatesCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return err
	}
	c.Request().ParseForm()
	data := strings.Join(c.Request().Form["categories"], ", ")
	new_template := models.Template{}
	new_template.Id = uuid.New()
	new_template.User = uuid.MustParse(user.Value)
	new_template.Categories = data
	if err := models.CreateTemplate(&new_template); err != nil {
		return err
	}

	return nil
}
