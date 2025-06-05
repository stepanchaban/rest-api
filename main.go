package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var task string

type RequestBody struct {
	Task string `json:"task"`
}

func postHandler(c echo.Context) error {
	var req RequestBody

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	task = req.Task

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Task updated",
	})
}

func getHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello, " + task)
}

func main() {
	e := echo.New()

	e.POST("/task", postHandler)
	e.GET("/task", getHandler)

	e.Start("localhost:8080")
}

