package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Todo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"is_done"`
}

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	e.GET("/", root)
	e.GET("/todos", todos)
	e.GET("/todos/:id", todo)

	// start server
	e.Start(":8080")
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "root page")
}

func todos(c echo.Context) error {
	return c.String(http.StatusOK, "todos page")
}

func todo(c echo.Context) error {
	todo := &Todo{
		ID:     100,
		Name:   "taka",
		IsDone: true,
	}
	return c.JSON(http.StatusOK, todo)
}
