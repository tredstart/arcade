package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	// should render a home page
	// also it is only accessible to logged in users
	_, err := utils.ReadCookie(c, "user")
	if err != nil {
		// probably redirect to login or smth
		return err
	}

	return views.HomePage().Render(c.Request().Context(), c.Response().Writer)
}

func LoginForm(c echo.Context) error {
	return views.Login().Render(c.Request().Context(), c.Response().Writer)
}

func Login(c echo.Context) error {
	username := c.Request().FormValue("username")

	// TODO: there should be hashing
	p := c.Request().FormValue("password")
	user, err := models.FetchUserByUsername(username)
	if err != nil {
		return err
	}
	if p != user.Password {
		return &utils.CustomError{S: "Wrong password"}
	}
	utils.WriteCookie(c, "name", user.Name)
	utils.WriteCookie(c, "user", user.Id.String())
	return nil
}

func RegisterForm(c echo.Context) error {
	return views.Register().Render(c.Request().Context(), c.Response().Writer)
}

func Register(c echo.Context) error {
	var user models.User
	password := c.Request().FormValue("password")
	if password != c.Request().FormValue("confirm") {
		return &utils.CustomError{S: "Passwords are not same"}
	}
	user.Id = uuid.New()
	user.Name = c.Request().FormValue("name")
	user.Username = c.Request().FormValue("username")
	// TODO: this should also be hashed
	user.Password = password
	if err := models.CreateUser(&user); err != nil {
		return err
	}
	// TODO: figure out how to do redirect
	//else {
	// 	id := "/profile/" + user.ID.String()
	// 	http.Redirect(w, r, id, http.StatusMovedPermanently)
	// }
	utils.WriteCookie(c, "name", user.Name)
	utils.WriteCookie(c, "user", user.Id.String())
	return nil
}

func LoginAsGuestForm(c echo.Context) error {
	return views.LoginAsGuest().Render(c.Request().Context(), c.Response().Writer)
}

func LoginAsGuest(c echo.Context) error {
	name := c.FormValue("name")
	utils.WriteCookie(c, "name", name)
	// Redirect to home?
	return nil
}

func History(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		// probably redirect to login or smth
		return err
	}
	retros, _ := models.FetchRetrosByUser(user.Value)
	return views.HistoryPage(retros).Render(c.Request().Context(), c.Response().Writer)
}
