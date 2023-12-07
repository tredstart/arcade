package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


func Home(c echo.Context) error {
    // should render a home page
    // also it is only accessible to logged in users
    username, err := utils.ReadCookie(c, "username")
    if err != nil {
        // probably redirect to login or smth
        return err
    }
    user, err := models.FetchUserByUsername(username.Value)
    if err != nil {
        return err
    }

    _, err = models.FetchLatestRetro(user.Id)
    if err != nil {
        return err
    }

    return nil
}

func LoginForm(c echo.Context) error {
    // should handle Login form 
    return nil
}

func Login(c echo.Context) error {
    // should validate input and put user in a cookie if success and redirect to hame
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
    utils.WriteCookie(c, "username", user.Username)
    return nil
}

func RegisterForm(c echo.Context) error {
    // should render RegisterForm
    return nil
}

func Register(c echo.Context) error {
    // should validate input and create new user
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
    utils.WriteCookie(c, "username", user.Username)
    return nil
}

func LoginAsGuestForm(c echo.Context) error {
    // when you just need to access as guest
    return nil
}

func LoginAsGuest(c echo.Context) error {
    name := c.FormValue("name")
    utils.WriteCookie(c, "name", name)
    // Redirect to home?
    return nil
}

func History(c echo.Context) error {
    // should render a history of retros for logged in user
    username, err := utils.ReadCookie(c, "username")
    if err != nil {
        // probably redirect to login or smth
        return err
    }
    user, err := models.FetchUserByUsername(username.Value)
    if err != nil {
        return err
    }

    _, err = models.FetchRetrosByUser(user.Id)
    if err != nil {
        return err
    }

    return nil
}
