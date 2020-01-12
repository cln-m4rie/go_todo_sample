package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = []Todo{
		{1, "test1"},
		{2, "test2"},
		{3, "test3"},
	}
	todoJSON = `[{"id":1,"name":"test1"},{"id":2,"name":"test2"},{"id":3,"name":"test3"}]`
)

func TestGetTodo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos")
	h := &Handler{mockDB}

	// Assertions
	if assert.NoError(t, h.ListTodo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, todoJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
