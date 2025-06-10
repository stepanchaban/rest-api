package main

import (
	"log"
	"my-api/internal/db"
	"my-api/internal/handlers"
	"my-api/internal/taskService"

	"github.com/labstack/echo/v4"
)



func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.POST("/tasks", taskHandlers.PostHandler)
	e.GET("/tasks", taskHandlers.GetHandler)
	e.PATCH("/tasks/:id", taskHandlers.PatchHandler)
	e.DELETE("/tasks/:id", taskHandlers.DeleteHandler)

	e.Start("localhost:8080")
}
