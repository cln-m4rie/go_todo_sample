package main

import (
	"github.com/cln-m4rie/go_todo_sample/handler"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/todos", handler.ListTodo)
	e.POST("/todos", handler.CreateTodo)
	e.GET("/todos/:id", handler.DetailTodo)
	e.Logger.Fatal(e.Start(":1323"))
}
