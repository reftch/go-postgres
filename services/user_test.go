package services

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/reftch/go-postgres/models"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*models.User{
		"1": &models.User{ID: 1, Name: "Jon Snow", Age: 32},
	}
	userJSON = `{"ID":1,Name:"Jon Snow","Age":"32"}`
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &UserService{nil}

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
