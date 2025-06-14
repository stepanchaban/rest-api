package main

import (
	"log"
	"my-api/internal/db"

	"my-api/internal/handlers"
	"my-api/internal/taskService"
	"my-api/internal/userService"
	"my-api/internal/web/tasks"
	"my-api/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database)
	tasksService := taskService.NewTaskService(tasksRepo)
	tasksHandlers := handlers.NewTaskHandler(tasksService)

	userRepo := userService.NewUserRepository(database)
	userService := userService.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasksStrictHandler := tasks.NewStrictHandler(tasksHandlers, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	usersStrictHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
