package taskService

type Task struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}

type RequestBody struct {
	Task string `json:"task"`
	IsDone bool `json:"is_done"`
}