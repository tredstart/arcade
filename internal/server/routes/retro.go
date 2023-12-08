package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RetroPage(c echo.Context) error {
	// display a form for every category in given retro
	return nil
}

func RetroItemCreate(c echo.Context) error {
	// handle data from an items on retro box page
	// get name of the category and content of new entry
	return nil
}

func RetroNewForm(c echo.Context) error {
	// not sure that this is a form lol
	return nil
}

func RetroCreate(c echo.Context) error {
	// handle a data from a retro form (or request)
	// it should use template and current user, and add a date of creation
	return nil
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
    data :=  strings.Join(c.Request().Form["categories"], ", ")
    new_template := models.Template{}
    new_template.Id = uuid.New()
    new_template.User = uuid.MustParse(user.Value)
    new_template.Categories = data
    if err := models.CreateTemplate(&new_template); err != nil {
        return err
    }

	return nil
}
