package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Saved(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	retros, _ := models.FetchSaved(user.Value)
	retros = utils.Reverse(retros)
	return views.SavedPage(retros).Render(c.Request().Context(), c.Response().Writer)
}

func DeleteSave(c echo.Context) error {
	_, err := utils.ReadCookie(c, "user")

	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("save")

	if err = models.DeleteSave(id); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return nil
}

func ButtonDelete(c echo.Context) error {
	_, err := utils.ReadCookie(c, "user")

	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("save")
	retro := c.Param("retro")

	if err = models.DeleteSave(id); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return views.SaveButton(retro, "").Render(c.Request().Context(), c.Response())
}

func CreateSave(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	c.Request().ParseForm()
	retro := c.FormValue("retro_id")
	id := uuid.NewString()
	if err := models.CreateSave(id, retro, user.Value); err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return views.SaveButton(retro, id).Render(c.Request().Context(), c.Response())

}

func VerifySave(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("retro")
	save, _ := models.FetchSaveByUserAndRetro(user.Value, id)
	log.Println("correct state: ", save == "")
	return views.SaveButton(id, save).Render(c.Request().Context(), c.Response())
}
