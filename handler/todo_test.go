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
	mockDB = map[string]*Todo{
		"1": &Todo{1, "test1"},
	}
	todoListJSON = `[{"id":1,"name":"test1"},{"id":2,"name":"test2"},{"id":3,"name":"test3"}]`
	todoJSON     = `{"id":1,"name":"test1"}`
)

func TestTodoListHandler_ListTodo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos")
	h := &TodoHandler{mockDB}

	// Assertions
	if assert.NoError(t, h.ListTodo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, todoListJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestTodoHandler_CreateTodo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos")
	h := &TodoHandler{mockDB}

	// Assertions
	if assert.NoError(t, h.CreateTodo(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, todoJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}

func TestTodoHandler_GetTodo(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(todoJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todos/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &TodoHandler{mockDB}

	// Assertions
	if assert.NoError(t, h.GetTodo(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, todoJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
