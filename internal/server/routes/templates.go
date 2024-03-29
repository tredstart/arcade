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
        return c.Redirect(http.StatusSeeOther, "/login")
	}
	return views.CreateTemplateForm("").Render(c.Request().Context(), c.Response().Writer)
}
func TemplatesCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
        return c.Redirect(http.StatusSeeOther, "/login")
	}
	c.Request().ParseForm()
	categories := c.Request().Form["categories"]
	data := strings.Join(categories, ", ")
	for _, char := range data {
		if unicode.IsDigit(char) {
			return views.CreateTemplateForm("Template category cannot contain numbers").Render(c.Request().Context(), c.Response().Writer)
		}
	}
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

func TemplatesDelete(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")

	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("template_id")

    if err = models.DeleteTemplate(id, user.Value); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
    }

    return nil
}
