package taskRepository

import (
	"tasks-app-api/pkg/domain/response"
	"tasks-app-api/pkg/domain/task"
	databaseHelpers "tasks-app-api/pkg/useCases/Helpers/databaseHelper"
)

type TaskRepository struct {
}

type Repository interface {
	CreateTask(task *task.Task) response.Status
	GetAllTask() ([]task.Task, response.Status)
	GetTask(taskId int) (task.Task, response.Status)
	UpdateTask(task task.Task) response.Status
}

func (tr *TaskRepository) CreateTask(task *task.Task) response.Status {
	db := databaseHelpers.GetDB()
	result := db.Create(&task)

	if result.Error != nil {
		return response.InternalServerError
	}
	return response.SuccessfulCreation
}

func (tr *TaskRepository) GetAllTask() ([]task.Task, response.Status) {
	db := databaseHelpers.GetDB()
	tasks := []task.Task{}
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, response.InternalServerError
	}
	return tasks, response.TaskFound
}

func (tr *TaskRepository) GetTask(taskId int) (task.Task, response.Status) {
	db := databaseHelpers.GetDB()
	task := task.Task{}
	result := db.Find(&task).Where("id = ?", taskId)
	if result.Error != nil {
		return task, response.InternalServerError
	}
	return task, response.TaskFound
}

func (tr *TaskRepository) UpdateTask(task task.Task) response.Status {
	db := databaseHelpers.GetDB()
	result := db.Save(&task)
	if result.Error != nil {
		return response.InternalServerError
	}
	return response.SuccessfulUpdate
}
