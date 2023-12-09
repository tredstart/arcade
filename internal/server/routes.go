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

	e.GET("/", s.IndexPage)
    e.GET("/login", routes.LoginForm) // -> also this can show form for guest(although I'm not sure about this)
    e.POST("/login", routes.Login)
    e.GET("/register", routes.RegisterForm)
    e.POST("/register", routes.Register)
    e.GET("/history", routes.History)

    // basically the whole retro logic that i can group into retro package
    e.GET("/retro/:id", routes.RetroPage) // -> should verify if user is in local storage or some shit
    e.POST("/retro/:id", routes.RetroItemCreate) // -> does a post to create an item in retro box
    e.POST("/retro/new", routes.RetroCreate)
    e.GET("/templates", routes.Templates)
    e.GET("/templates/new", routes.TemplatesNew)
    e.POST("/templates/new", routes.TempalatesCreate)

	return e
}

func (s *Server) IndexPage(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/history")
}
