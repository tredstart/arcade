package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"net/http"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	return views.CreateTemplateForm("").Render(c.Request().Context(), c.Response().Writer)
}

func TemplatesCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	c.Request().ParseForm()
	data := strings.Join(c.Request().Form["categories"], ", ")
    for _, category := range data {
        if unicode.IsDigit(category) {
			return views.CreateTemplateForm("Categories cannot be entirely nubmers").Render(c.Request().Context(), c.Response().Writer)
		}
    }
	new_template := models.Template{}
	new_template.Id = uuid.New()
	new_template.User = uuid.MustParse(user.Value)
	new_template.Categories = data
	if err := models.CreateTemplate(&new_template); err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}

	return c.Redirect(http.StatusSeeOther, "/templates")
}
