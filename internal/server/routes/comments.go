package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func CommentsView(c echo.Context) error {

	if _, err := utils.ReadCookie(c, "name"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	record_id := c.Param("record_id")
	comments, _ := models.FetchCommentsByRecord(record_id)
	return views.CommentsBlock(comments, record_id).Render(c.Request().Context(), c.Response().Writer)
}

func CommentsCount(c echo.Context) error {

	if _, err := utils.ReadCookie(c, "name"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	record_id := c.Param("record_id")
	count, err := models.CountComments(record_id)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return c.String(http.StatusOK, fmt.Sprint(count))

}

func CommentsAdd(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	var comment models.Comment
	comment.Id = uuid.New()
	comment.Author = name.Value
	comment.Content = c.FormValue("content")
	record_id := c.Param("record_id")
	comment.Record = uuid.MustParse(record_id)
	if err = models.CreateComment(&comment); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return views.Comment(comment).Render(c.Request().Context(), c.Response().Writer)
}

func CommentLike(c echo.Context) error {

	if _, err := utils.ReadCookie(c, "name"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	var liked string

	if l, err := utils.ReadCookie(c, "comments"); err == nil {
		liked = l.Value
	}

	comment_id := c.Param("comment_id")
	likes, err := strconv.Atoi(c.FormValue("likes"))

	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	if strings.Contains(liked, comment_id) {
		return views.EmptyLike(likes).Render(c.Request().Context(), c.Response().Writer)
	}

	likes++
	err = models.LikeTheComment(comment_id, likes)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

    liked += fmt.Sprintf(" %s", comment_id)

    utils.WriteCookie(c, "comments", liked)

	return views.EmptyLike(likes).Render(c.Request().Context(), c.Response().Writer)
}
