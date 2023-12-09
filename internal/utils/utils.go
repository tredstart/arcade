package utils

import (
	"log"
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
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    log.Println(err)
    return err == nil
}
