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
	e.GET("/retros", routes.History)
	e.GET("/guest", routes.LoginAsGuestForm)
	e.POST("/guest", routes.LoginAsGuest)

	e.GET("/profile", routes.UpdateUserForm)
	e.PUT("/profile", routes.UpdateUser)

	e.GET("/retro/:id", routes.RetroPage)
	e.GET("/retro/:id/record/:record_id", routes.RecordUpdateForm)
	e.PATCH("/retro/:id/record/:record_id", routes.RetroItemUpdate)
	e.DELETE("/retro/:id/record/:record_id", routes.RetroItemDelete)
	e.POST("/retro/:id", routes.RetroItemCreate)
	e.DELETE("/retro/:id", routes.RetroDelete)
	e.POST("/retro/new", routes.RetroCreate)
	e.GET("/templates", routes.Templates)
	e.DELETE("/templates/:template_id", routes.TemplatesDelete)
	e.GET("/templates/new", routes.TemplatesNew)
	e.POST("/templates/new", routes.TemplatesCreate)

	e.POST("/retro/:id/change-visibility", routes.RetroMakeVisible)
	e.PATCH("/record/:record_id", routes.RecordLike)
	e.GET("/record/:record_id", routes.RecordView)
	e.GET("/record/:record_id/likes/:likes", routes.RecordLikes)
	e.GET("/record/:record_id/comments", routes.CommentsView)
	e.POST("/record/:record_id/comments", routes.CommentsAdd)
	e.PATCH("/comments/:comment_id", routes.CommentLike)
	e.GET("/comments/:comment_id/:likes", routes.CommentsLike)
	e.GET("/count-comments/:record_id", routes.CommentsCount)

	e.GET("/saved", routes.Saved)
	e.DELETE("/saved/:save", routes.DeleteSave)
	e.DELETE("/button-delete/:retro/:save", routes.ButtonDelete)
	e.POST("/saved", routes.CreateSave)
	e.GET("/verify-save/:retro", routes.VerifySave)

	return e
}

func (s *Server) IndexPage(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/retros")
}
