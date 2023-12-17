package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CommentsView(c echo.Context) error {

    if _, err := utils.ReadCookie(c, "name"); err != nil {
        return c.Redirect(http.StatusSeeOther, "/guest")
	}

    record_id := c.Param("id")
    _, _ = models.FetchCommentsByRecord(record_id)
    return nil // some form here
}

func CommentsAdd(c echo.Context) error {
    name, err := utils.ReadCookie(c, "name");
    if err != nil {
        return c.Redirect(http.StatusSeeOther, "/guest")
	}

    var comment models.Comment
    comment.Id = uuid.New()
    comment.Author = name.Value
    comment.Content = c.FormValue("content")
    record_id := c.Param("id")
    comment.Record = uuid.MustParse(record_id)
    if err = models.CreateComment(&comment); err != nil {
        return err
    }

    return nil
}

func CommentLike(c echo.Context) error {

    if _, err := utils.ReadCookie(c, "name"); err != nil {
        return c.Redirect(http.StatusSeeOther, "/guest")
	}
    return nil
}
