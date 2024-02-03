package routes

import (
	"arcade/internal/models"
	"arcade/internal/utils"
	"arcade/internal/views"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RetroPage(c echo.Context) error {

	retro_id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	name, err := utils.ReadCookie(c, "name")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/guest?next=/retro/"+retro_id.String())
	}
	retro, err := models.FetchRetro(retro_id)
	if err != nil {
        log.Error("retro: ", err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	template, err := models.FetchTemplate(retro.Template)
	if err != nil {
        log.Error("template: ", err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	categories := strings.Split(template.Categories, ", ")
	context := make(map[string][]models.Record)
	c_ids := make(map[string]string)
	auth_table := make(map[string]bool)

	for _, category := range categories {
		context[category] = []models.Record{}

		regex := regexp.MustCompile(`[^a-zA-Z_\-]`)
		c_ids[category] = regex.ReplaceAllString(category, "_")
		if retro.Visible {
			records, _ := models.FetchRecordsByRetro(retro_id)
			for _, record := range records {
				if record.Category == category {
					context[category] = append(context[category], record)
				}
				auth_table[record.Id.String()] = name.Value == record.Author
			}
		} else {
			records, _ := models.FetchRecordsByRetroAndName(retro_id, name.Value)
			for _, record := range records {
				if record.Category == category {
					context[category] = append(context[category], record)
				}
				auth_table[record.Id.String()] = name.Value == record.Author
			}
		}

	}

	user_id, err := utils.ReadCookie(c, "user")
	authorized := false

	if err == nil {
		if u, e := models.FetchUser(user_id.Value); e == nil && u.Id == retro.User {
			authorized = true
		}
	}

	return views.RetroPage(context, c_ids, "Retro from "+string(retro.Created), retro, authorized, auth_table).Render(c.Request().Context(), c.Response().Writer)
}

func RecordLike(c echo.Context) error {

	if _, err := utils.ReadCookie(c, "name"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	var liked string

	if l, err := utils.ReadCookie(c, "records"); err == nil {
		liked = l.Value
	}

	c.Request().ParseForm()

	record_id := c.Param("record_id")
	likes, err := strconv.Atoi(c.FormValue("likes"))
	active := "active"

	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	if strings.Contains(liked, record_id) {
		likes--
		liked = strings.ReplaceAll(liked, fmt.Sprintf(" %s", record_id), "")
		active = ""
	} else {
		likes++
		liked += fmt.Sprintf(" %s", record_id)
	}

	err = models.LikeTheRecord(record_id, likes)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	utils.WriteCookie(c, "records", liked)

	return views.Like(fmt.Sprint(likes), fmt.Sprintf("/record/%s", record_id), active).Render(c.Request().Context(), c.Response().Writer)
}

func RecordLikes(c echo.Context) error {

	if _, err := utils.ReadCookie(c, "name"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	var liked string

	if l, err := utils.ReadCookie(c, "records"); err == nil {
		liked = l.Value
	}

	record_id := c.Param("record_id")
	likes := c.Param("likes")
	var active string
	if strings.Contains(liked, record_id) {
		active = "active"
	}
	return views.Like(fmt.Sprint(likes), fmt.Sprintf("/record/%s", record_id), active).Render(c.Request().Context(), c.Response().Writer)
}

func RecordView(c echo.Context) error {

	record_id := c.Param("record_id")

	record, err := models.FetchRecord(record_id)

	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

    log.Warn(c.Request().Header)
	if c.Request().Header["Load-Bottom"] != nil {
		return views.RetroItemBottom(record).Render(c.Request().Context(), c.Response().Writer)
	} else {
		return views.RetroItem(record, true, record.Retro.String()).Render(c.Request().Context(), c.Response().Writer)
	}

}

func RetroMakeVisible(c echo.Context) error {
	if _, err := utils.ReadCookie(c, "user"); err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	retro_id := c.Param("id")
	visible := c.FormValue("visible")
	visibility := !(visible == "true")

	if err := models.RetroSetVisibility(retro_id, visibility); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return c.Redirect(http.StatusSeeOther, "/retro/"+retro_id)
}

func RetroItemCreate(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")
	id := c.Param("id")
	retro_id := uuid.MustParse(id)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/guest?next=/retro/"+retro_id.String())
	}
	category := c.FormValue("category")
	content := c.FormValue("content")
	var record models.Record
	record.Id = uuid.New()
	record.Retro = retro_id
	record.Author = name.Value
	record.Category = category
	record.Content = content
	if err = models.CreateRecord(&record); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return views.RetroItem(record, true, retro_id.String()).Render(c.Request().Context(), c.Response().Writer)
}

func RetroCreate(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	c.Request().ParseForm()
	template_id := c.FormValue("template_id")
	new_retro := models.Retro{}
	new_retro.Id = uuid.New()
	new_retro.User = uuid.MustParse(user.Value)
	new_retro.Template = uuid.MustParse(template_id)
	new_retro.Created = time.Now().Format("2006-01-02")
	new_retro.Visible = false
	if err := models.CreateRetro(&new_retro); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}
	return c.Redirect(http.StatusSeeOther, "/retro/"+new_retro.Id.String())
}

func RetroDelete(c echo.Context) error {
	user, err := utils.ReadCookie(c, "user")

	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	id := c.Param("id")

	if err = models.RetroDelete(id, user.Value); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return nil
}

func RetroItemDelete(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")

	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	retro_id := c.Param("id")
	record_id := c.Param("record_id")

	if err = models.RecordDelete(record_id, name.Value, retro_id); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return nil
}

func RecordUpdateForm(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")

	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	retro_id := c.Param("id")
	record_id := c.Param("record_id")

	content, err := models.FetchRecordContent(record_id, retro_id, name.Value)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return views.UpdateRetroItemForm(content, retro_id, record_id).Render(c.Request().Context(), c.Response().Writer)
}

func RetroItemUpdate(c echo.Context) error {
	name, err := utils.ReadCookie(c, "name")

	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	retro_id := c.Param("id")
	record_id := c.Param("record_id")
	content := c.FormValue("content")

	if err = models.RecordUpdateContent(record_id, content, name.Value, retro_id); err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	record, err := models.FetchRecord(record_id)
	if err != nil {
		log.Error(err.Error())
		return c.String(http.StatusInternalServerError, "Oops, something went wrong. Please try again")
	}

	return views.RetroItem(record, true, retro_id).Render(c.Request().Context(), c.Response().Writer)
}
