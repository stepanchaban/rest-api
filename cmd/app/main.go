package main

import (
	"log"
	"my-api/internal/db"
	"my-api/internal/handlers"
	"my-api/internal/taskService"
	"my-api/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func main() {
	database, err := db.InitDB()
	
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)
	

	// e.POST("/tasks", taskHandlers.PostTasks)
	// e.GET("/tasks", taskHandlers.GetTasks)
	// e.PATCH("/tasks/:id", taskHandlers.PatchHandler)
	// e.DELETE("/tasks/:id", taskHandlers.DeleteHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
