package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-wms-poc/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mockDB = map[int]*models.Area{
		1: {Name: "Inbound"},
	}
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

func TestFindAreaById(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/areas/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &Handler{}

	// Assertions
	if assert.NoError(t, h.FindAreas(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
