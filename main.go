package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

type RequestBody struct {
	Task string `json:"task"`
}

var tasks []Task

func postHandler(c echo.Context) error {
	var req RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	newTask := Task{
		ID:   uuid.NewString(),
		Task: req.Task,
	}

	tasks = append(tasks, newTask)

	return c.JSON(http.StatusCreated, newTask)
}

func getHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func patchHandler(c echo.Context) error {
	id := c.Param("id")

	var req RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	for i, t := range tasks {
		if t.ID == id {
			if req.Task != "" {
				tasks[i].Task = req.Task
			}
			return c.JSON(http.StatusOK, tasks[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Task not found",
	})
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Task not found",
	})
}

func main() {
	e := echo.New()

	e.POST("/task", postHandler)
	e.GET("/task", getHandler)
	e.PATCH("/task/:id", patchHandler)
	e.DELETE("/task/:id", deleteHandler)

	e.Start("localhost:8080")
}
