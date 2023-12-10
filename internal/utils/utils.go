package utils

import (
	"net/http"
	"os"
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
    cookie.SameSite = http.SameSiteStrictMode
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
    return err == nil
}

func ReadEnvVar(name string) (string, error) {
    env_var, exists := os.LookupEnv(name)
    if !exists || env_var == "" {
        err := &CustomError{S: "Variable: " + name + "is empty or does not exists"}
        return "", err
    }
    return env_var, nil
}
