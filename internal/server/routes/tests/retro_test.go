package tests

import (
	"arcade/internal/database"
	"arcade/internal/models"
	"arcade/internal/server/routes"
	"arcade/internal/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplatesPageRedirects(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/templates", nil)

	c := e.NewContext(req, rec)

	assert.NoError(t, routes.Templates(c))
	require.Equal(t, http.StatusSeeOther, rec.Code)

}

func TestRetroPageRedirectsWithoutCookieName(t *testing.T) {
	template := createTestTemplate()
	retro := createTestRetro(*template)
	e := echo.New()
	rec := httptest.NewRecorder()
    rq := "/retro/"

	req := httptest.NewRequest(http.MethodGet, rq, nil)

	c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues(retro.Id.String())

	assert.NoError(t, routes.RetroPage(c))
	require.Equal(t, http.StatusSeeOther, rec.Code)
}

func createTestUser() *models.User {
	var user models.User
	user.Id = uuid.New()
	user.Username = "test_user-" + user.Id.String()
	user.Name = "Test User"
	password, _ := utils.HashPassword("testing")
	user.Password = string(password)
	if err := models.CreateUser(&user); err != nil {
		log.Fatalln(err)
	}
	return &user
}

func createTestTemplate() *models.Template {
	var template models.Template
	template.Id = uuid.New()
	template.Categories = "test, random, try me"
	template.User = createTestUser().Id
	if err := models.CreateTemplate(&template); err != nil {
		log.Fatalln(err)
	}
	return &template
}

func createTestRetro(template models.Template) *models.Retro {
	var retro models.Retro
	retro.Id = uuid.New()
	retro.Created = time.Now().Format("2006-01-02")
	retro.User = template.User
	retro.Template = template.Id
	if err := models.CreateRetro(&retro); err != nil {
		log.Fatalln(err)
	}
	return &retro
}

func createTestRecord(template models.Template, retro models.Retro) *models.Record {
	var record models.Record
	record.Id = uuid.New()
	record.Author = "Test User"
	record.Retro = retro.Id
	record.Content = "test content present"
	record.Category = strings.Split(template.Categories, ", ")[0]
	if err := models.CreateRecord(&record); err != nil {
		log.Fatalln(err)
	}
	return &record
}

func eraseTestData(db *sqlx.DB) {
	_, err := db.Exec("DELETE FROM user")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM template")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM retro")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("DELETE FROM record")
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	// TODO find a way to fix this
	database.DB = sqlx.MustConnect("sqlite3", "../../../../testing.db")
	_ = m.Run()
	// teardown
	eraseTestData(database.DB)
}
