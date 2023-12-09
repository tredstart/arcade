package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomError struct {
	S string
}

func (e *CustomError) Error() string {
	return e.S
}

func WriteCookie(c echo.Context, name, value string) {
    cookie := new(http.Cookie)
    cookie.Name = name
    cookie.Value = value
    cookie.Expires = time.Now().Add(24 * time.Hour)
    c.SetCookie(cookie)
}

func ReadCookie(c echo.Context, name string) (*http.Cookie, error){
    cookie, err := c.Cookie(name)
    if err != nil {
        return nil, err
    }
    return cookie, nil
}

func HashPassword(password string) (string, error) {
	// Generate a hashed version of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
