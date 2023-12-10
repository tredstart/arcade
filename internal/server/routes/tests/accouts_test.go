package tests

import (
	"arcade/internal/server/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginPage(t *testing.T) {
    e := echo.New()
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/login", nil)

    c := e.NewContext(req, rec)

    assert.NoError(t, routes.LoginForm(c))
    require.Equal(t, http.StatusOK, rec.Code)
}

func TestRegisterPage(t *testing.T) {
    e := echo.New()
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/register", nil)

    c := e.NewContext(req, rec)

    assert.NoError(t, routes.RegisterForm(c))
    require.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginAsGuest(t *testing.T) {
    e := echo.New()
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/guest", nil)

    c := e.NewContext(req, rec)

    assert.NoError(t, routes.LoginAsGuestForm(c))
    require.Equal(t, http.StatusOK, rec.Code)
}

func TestHistoryPageRedirects(t *testing.T) {
    e := echo.New()
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/history", nil)

    c := e.NewContext(req, rec)

    assert.NoError(t, routes.History(c))
    require.Equal(t, http.StatusSeeOther, rec.Code)

}

