package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB

func initDB() {
	dsn :="host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
}

type Task struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}

type RequestBody struct {
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}

func getHandler(c echo.Context) error {
	var task []Task

	if err := db.Find(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, task)
}

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
		IsDone: req.IsDone,
	}

	if err := db.Create(&newTask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, newTask)
}



func patchHandler(c echo.Context) error {
	id := c.Param("id")

	var req RequestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request",
		})
	}

	var task Task

	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not find task"})
	}

	task.Task = req.Task
	task.IsDone = req.IsDone

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {
	initDB()
	e := echo.New()

	e.POST("/tasks", postHandler)
	e.GET("/tasks", getHandler)
	e.PATCH("/tasks/:id", patchHandler)
	e.DELETE("/tasks/:id", deleteHandler)

	e.Start("localhost:8080")
}
