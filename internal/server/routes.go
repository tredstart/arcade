package server

import (
	"arcade/internal/server/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Static("/static", "assets")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.IndexPage)
	e.GET("/login", routes.LoginForm)
	e.POST("/login", routes.Login)
	e.GET("/register", routes.RegisterForm)
	e.POST("/register", routes.Register)
	e.GET("/history", routes.History)
	e.GET("/guest", routes.LoginAsGuestForm)
	e.POST("/guest", routes.LoginAsGuest)

    e.GET("/profile", routes.UpdateUserForm)
    e.PUT("/profile", routes.UpdateUser)

	e.GET("/retro/:id", routes.RetroPage)
	e.POST("/retro/:id", routes.RetroItemCreate)
	e.POST("/retro/new", routes.RetroCreate)
	e.GET("/templates", routes.Templates)
	e.GET("/templates/new", routes.TemplatesNew)
	e.POST("/templates/new", routes.TemplatesCreate)

    e.POST("/retro/:id/change-visibility", routes.RetroMakeVisible)
    e.PATCH("/record/:id", routes.RecordLike)
    e.GET("/record/:id/comments", routes.CommentsView)
    e.POST("/record/:id/comments", routes.CommentsAdd)
    e.PATCH("record/:id/comments", routes.CommentLike)
    e.GET("/count-comments/:retro-id", routes.CommentsCount)

	return e
}

func (s *Server) IndexPage(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/history")
}
