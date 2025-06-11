package taskService

import "github.com/oapi-codegen/runtime/types"

type Task struct {
	ID   types.UUID `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}

type RequestBody struct {
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}