package routes

import "github.com/labstack/echo/v4"


func Home(c echo.Context) error {
    // should render a home page
    // also it is only accessible to logged in users
    return nil
}

func LoginForm(c echo.Context) error {
    // should handle Login form 
    return nil
}

func Login(c echo.Context) error {
    // should validate input and put user in a cookie if success and redirect to hame
    return nil
}

func RegisterForm(c echo.Context) error {
    // should render RegisterForm
    return nil
}

func Register(c echo.Context) error {
    // should validate input and create new user
    return nil
}

func LoginAsGuestForm(c echo.Context) error {
    // when you just need to access as guest

    // should take in name and put it in cookie
    // also probably only accessible when a there is no user in cookie while trying to open a retro page
    return nil
}

func History(c echo.Context) error {
    // should render a history of retros for logged in user
    return nil
}
