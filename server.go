package main

import (
	"github.com/cln-m4rie/go_todo_sample/handler"
	"net/http"

	"github.com/labstack/echo"
)

var (
	mockDB = []handler.Todo{
		{1, "test1"},
		{2, "test2"},
		{3, "test3"},
	}
)

func main() {
	e := echo.New()
	h := handler.Handler{mockDB}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/todos", h.ListTodo)
	e.POST("/todos", handler.CreateTodo)
	e.GET("/todos/:id", handler.GetTodo)
	e.PUT("/todos/:id", handler.UpdateTodo)
	e.Logger.Fatal(e.Start(":1323"))
}
