package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

func LoginForm(c echo.Context) error {
	return views.Login("").Render(c.Request().Context(), c.Response().Writer)
}

func Login(c echo.Context) error {
	username := c.Request().FormValue("username")
	if len(username) < 3 {
		return views.Login("Username must be longer than 3 characters").Render(c.Request().Context(), c.Response().Writer)
	}

	p := c.Request().FormValue("password")
	user, err := models.FetchUserByUsername(username)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return views.Login("No user found with this username").Render(c.Request().Context(), c.Response().Writer)
		}
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again.")
	}
	if !utils.CheckPasswordHash(p, user.Password) {
		return views.Login("Wrong password").Render(c.Request().Context(), c.Response().Writer)
	}
	utils.WriteCookie(c, "name", user.Name)
	utils.WriteCookie(c, "user", user.Id.String())
	return c.Redirect(http.StatusSeeOther, "/templates")
}

func RegisterForm(c echo.Context) error {
	return views.Register("").Render(c.Request().Context(), c.Response().Writer)
}

func Register(c echo.Context) error {
	var user models.User
	password := c.Request().FormValue("password")

	if len(password) < 6 {
		return views.Register("Password cannot be less than 6 characters").Render(c.Request().Context(), c.Response().Writer)
	}
	if password != c.Request().FormValue("confirm") {
		err := &utils.CustomError{S: "Passwords are not the same"}
		return views.Register(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	user.Id = uuid.New()
	user.Name = c.Request().FormValue("name")
	if len(user.Name) == 0 {
		return views.Register("Name cannot be empty").Render(c.Request().Context(), c.Response().Writer)
	}
	user.Username = c.Request().FormValue("username")
	if len(user.Username) < 3 {
		return views.Register("Username cannot be less than 3 characters").Render(c.Request().Context(), c.Response().Writer)
	}
	p, err := utils.HashPassword(password)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	user.Password = p
	if err := models.CreateUser(&user); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	utils.WriteCookie(c, "name", user.Name)
	utils.WriteCookie(c, "user", user.Id.String())
	return c.Redirect(http.StatusSeeOther, "/templates")
}

func UpdateUserForm(c echo.Context) error {

	user_id, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	user, err := models.FetchUser(user_id.Value)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return views.UpdateUser(user, "").Render(c.Request().Context(), c.Response().Writer)
}

func UpdateUser(c echo.Context) error {

	user_id, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	user, err := models.FetchUser(user_id.Value)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	password := c.Request().FormValue("password")
	if password != c.Request().FormValue("confirm") {
		err := &utils.CustomError{S: "Passwords are not the same"}
		return views.UpdateUser(user, err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}

	// TODO: fix validation (move it to func)
	user.Name = c.Request().FormValue("name")
	user.Username = c.Request().FormValue("username")
	p, err := utils.HashPassword(password)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	user.Password = p
	if err := models.UpdateUser(&user); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return c.Redirect(http.StatusSeeOther, "/templates")
}

func LoginAsGuestForm(c echo.Context) error {
	return views.LoginAsGuest().Render(c.Request().Context(), c.Response().Writer)
}

func LoginAsGuest(c echo.Context) error {
	name := c.FormValue("name")
	next := c.QueryParams()["next"][0]
	if next == "" {
		next = "/"
	}
	utils.WriteCookie(c, "name", name)
	return c.Redirect(http.StatusSeeOther, next)
}

func History(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	retros, _ := models.FetchRetrosByUser(user.Value)
	retros = utils.Reverse[models.Retro](retros)
	return views.HistoryPage(retros).Render(c.Request().Context(), c.Response().Writer)
}
