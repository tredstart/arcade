package server

import (
	"arcade/internal/server/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)
    // and this should be account thin idk
    // also how do I do "save to the local storage and shit?"
    e.GET("/login", routes.LoginForm) // -> also this can show form for guest(although I'm not sure about this)
    e.POST("/login", routes.Login)
    e.GET("/signup", routes.RegisterForm)
    e.POST("/signup", routes.Register)
    e.GET("/home", routes.Home) // -> should probably show the latest retro
    e.GET("/history", routes.History)

    // basically the whole retro logic that i can group into retro package
    // e.GET("/retro/:url") -> should verify if user is in local storage or some shit
    // e.POST("/retro/:url") -> does a post to create an item in retro box
    // e.POST("/retro/new")
    e.GET("/retro/templates", routes.Templates)
    e.GET("/retro/templates/new", routes.TemplatesNew)
    e.POST("/retro/templates/new", routes.TempalatesCreate)

	return e
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
