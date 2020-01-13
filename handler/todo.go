package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type (
	Todo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	TodoListHandler struct {
		Db []Todo
	}
	TodoHandler struct {
		Db Todo
	}
)

var (
	Todos = map[int]*Todo{}
	seq   = 1
)

func (h TodoListHandler) ListTodo(c echo.Context) error {
	todos := h.Db
	return c.JSON(http.StatusOK, todos)
}

func (h TodoHandler) CreateTodo(c echo.Context) error {
	todo := &Todo{
		Id: seq,
	}
	if err := c.Bind(todo); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, todo)
}

func (h TodoHandler)GetTodo(c echo.Context) error {
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
