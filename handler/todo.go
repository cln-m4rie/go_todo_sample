package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func ListTodo(c echo.Context) error {
	todos := []Todo{
		{Id: 1, Name: "test1"},
		{Id: 2, Name: "test2"},
	}
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, todo)
}

func DetailTodo(c echo.Context) error {
	var id int
	id, _ = strconv.Atoi(c.Param("id"))
	todo := Todo{
		Id: id,
		Name: "test1",
	}
	return c.JSON(http.StatusOK, todo)
}