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
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}

	return views.HomePage().Render(c.Request().Context(), c.Response().Writer)
}

func LoginForm(c echo.Context) error {
	return views.Login().Render(c.Request().Context(), c.Response().Writer)
}

func Login(c echo.Context) error {
	username := c.Request().FormValue("username")
	if len(username) < 3 {
		return views.ErrorBlock("Username cannot be less than 4 characters").Render(c.Request().Context(), c.Response().Writer)
	}

	// TODO: there should be hashing
	p := c.Request().FormValue("password")
	user, err := models.FetchUserByUsername(username)
	if err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	if p, err = utils.HashPassword(p); err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	if p != user.Password {
		err = &utils.CustomError{S: "Wrong password"}
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
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

	if len(password) < 6 {
		return views.ErrorBlock("Password cannot be less than 6 characters").Render(c.Request().Context(), c.Response().Writer)
	}
	if password != c.Request().FormValue("confirm") {
		err := &utils.CustomError{S: "Passwords are not same"}
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	user.Id = uuid.New()
	user.Name = c.Request().FormValue("name")
	if len(user.Name) == 0 {
		return views.ErrorBlock("Name cannot be empty").Render(c.Request().Context(), c.Response().Writer)
	}
	user.Username = c.Request().FormValue("username")
	if len(user.Username) < 3 {
		return views.ErrorBlock("Username cannot be less than 4 characters").Render(c.Request().Context(), c.Response().Writer)
	}
	user.Username = c.Request().FormValue("username")
	// TODO: this should also be hashed
	p, err := utils.HashPassword(password)
	if err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	user.Password = p
	if err := models.CreateUser(&user); err != nil {
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
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
	if len(name) == 0 {
		return views.ErrorBlock("Name cannot be empty").Render(c.Request().Context(), c.Response().Writer)
	}
	utils.WriteCookie(c, "name", name)
	// Redirect to home?
	return nil
}

func History(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		// probably redirect to login or smth
		return views.ErrorBlock(err.Error()).Render(c.Request().Context(), c.Response().Writer)
	}
	retros, _ := models.FetchRetrosByUser(user.Value)
	return views.HistoryPage(retros).Render(c.Request().Context(), c.Response().Writer)
}
