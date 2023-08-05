package taskHandler

import (
	"tasks-app-api/internal/data/infrastructure/taskRepository"
	"tasks-app-api/pkg/domain/response"
	"tasks-app-api/pkg/domain/task"
)

type TaskHandler struct {
	Repository taskRepository.Repository
}

type Handler interface {
	CreateTask(task task.Task) response.Status
	GetTask(taskId int) (interface{}, response.Status)
	GetAllTask() (interface{}, response.Status)
	UpdateTask(task task.Task) response.Status
}

func (th *TaskHandler) CreateTask(task task.Task) response.Status {
	return th.Repository.CreateTask(&task)
}

func (th *TaskHandler) GetTask(taskId int) (interface{}, response.Status) {
	return th.Repository.GetTask(taskId)
}

func (th *TaskHandler) GetAllTask() (interface{}, response.Status) {
	return th.Repository.GetAllTask()
}

func (th *TaskHandler) UpdateTask(task task.Task) response.Status {
	return th.Repository.UpdateTask(task)
}
