package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var (
	Todos = map[int]*Todo{}
	seq   = 1
)

func ListTodo(c echo.Context) error {
	todos := []Todo{
		{Id: 1, Name: "test1"},
		{Id: 2, Name: "test2"},
	}
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	todo := &Todo{
		Id: seq,
	}
	if err := c.Bind(todo); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, todo)
}

func GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := Todo{
		Id:   id,
		Name: "test1",
	}
	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	reqTodo := new(Todo)
	if err := c.Bind(reqTodo); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	todo := Todo{
		Id:   id,
		Name: reqTodo.Name,
	}
	return c.JSON(http.StatusOK, todo)
}
