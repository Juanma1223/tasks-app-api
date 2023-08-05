package taskRepository

import (
	"fmt"
	"tasks-app-api/pkg/domain/response"
	"tasks-app-api/pkg/domain/task"
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
	stmt, err := txn.Prepare(`INSERT INTO
                                task(
                                    id,
text,
list_id
                                )
                            VALUES
                                (?,?,?)`)
	if err != nil {
		return response.DBQueryError
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Printf("Error closing statement: %v", err)
		}
	}()
	row, err := stmt.Exec(
		task.Id,
		task.Text,
		task.ListId,
	)
	if err != nil {
		return response.DBExecutionError
	}

	id, err := row.LastInsertId()
	if err != nil {
		return response.DBLastRowIdError
	}
	task.Id = int(id)
	if err != nil {
		return response.DBExecutionError
	}

	return response.SuccessfulCreation
}

func (tr *TaskRepository) GetAllTask() ([]task.Task, response.Status) {
	stmt, err := txn.Prepare(`SELECT
                                id,
text,
list_id
                            FROM
                                task`)
	if err != nil {
		return nil, response.DBQueryError
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Printf("Error closing statement: %v", err)
		}
	}()
	res, err := stmt.Query()
	if err != nil {
		return nil, response.DBExecutionError
	}
	if res.Err() != nil {
		return nil, response.DBExecutionError
	}
	defer func() {
		if err := res.Close(); err != nil {
			fmt.Printf("Error closing statement: %v", err)
		}
	}()
	var tasks []task.Task
	for res.Next() {
		var task task.Task
		err = res.Scan(
			&task.Id,
			&task.Text,
			&task.ListId,
		)
		if err != nil {
			return nil, response.InternalServerError
		}
		tasks = append(tasks, task)
	}
	return tasks, response.TaskFound
}

func (tr *TaskRepository) GetTask(taskId int) (task.Task, response.Status) {
	var task task.Task
	stmt, err := txn.Prepare(`SELECT
                                id,
text,
list_id
                            FROM
                                task
                            WHERE
                                id = ?`)
	if err != nil {
		return task, response.DBQueryError
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Printf("Error closing statement: %v", err)
		}
	}()
	err = stmt.QueryRow(taskId).Scan(
		&task.Id,
		&task.Text,
		&task.ListId,
	)
	if err != nil {
		return task, response.DBExecutionError
	}
	return task, response.TaskFound
}

func (tr *TaskRepository) UpdateTask(task task.Task) response.Status {
	stmt, err := txn.Prepare(`UPDATE 
                                task
                            SET
                                id = ?,
text = ?,
list_id = ?
                            WHERE
                                id = ?`)
	if err != nil {
		return response.DBQueryError
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			fmt.Printf("Error closing statement: %v", err)
		}
	}()
	_, err = stmt.Exec(
		task.Id,
		task.Text,
		task.ListId,
	)
	if err != nil {
		return response.DBExecutionError
	}
	return response.SuccessfulUpdate
}
