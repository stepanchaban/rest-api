package taskService

import (
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type TaskService interface {
	CreateTask(task string, isDone bool) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTask(id string, task string, isDone bool) (Task, error)
	DeleteTask(id string) error
	GetTaskByID(id string) (Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}


func (s *taskService) CreateTask(task string, isDone bool) (Task, error) {

	id := types.UUID(uuid.New())

	newTask := Task{
		ID:     id,
		Task:   task,
		IsDone: isDone,
	}

	if err := s.repo.CreateTask(newTask); err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}


func (s *taskService) GetTaskByID(id string) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id string, task string, isDone bool) (Task, error) {
	existingTask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	existingTask.Task = task
	existingTask.IsDone = isDone

	if err := s.repo.UpdateTask(existingTask); err != nil {
		return Task{}, err
	}

	return existingTask, nil
}


func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}



