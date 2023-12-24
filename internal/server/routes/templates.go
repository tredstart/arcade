package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Templates(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	templates, _ := models.FetchTemplatesByUser(user.Value)
	templates = utils.Reverse[models.Template](templates)
	return views.Templates(templates).Render(c.Request().Context(), c.Response().Writer)
}

func TemplatesNew(c echo.Context) error {
	_, err := utils.ReadCookie(c, "user")
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return views.CreateTemplateForm("").Render(c.Request().Context(), c.Response().Writer)
}
func TemplatesCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	c.Request().ParseForm()
	categories := c.Request().Form["categories"]
	for _, category := range categories {
		if utils.IsStringNumeric(category) {
			return views.CreateTemplateForm("Categories cannot be entirely nubmers").Render(c.Request().Context(), c.Response().Writer)
		}
	}
	data := strings.Join(categories, ", ")
	new_template := models.Template{}
	new_template.Id = uuid.New()
	new_template.User = uuid.MustParse(user.Value)
	new_template.Categories = data
	if err := models.CreateTemplate(&new_template); err != nil {
        log.Error(err.Error())
        return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return c.Redirect(http.StatusSeeOther, "/templates")
}
