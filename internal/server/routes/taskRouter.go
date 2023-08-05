package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tasks-app-api/pkg/domain/task"
	"tasks-app-api/pkg/useCases/Handlers/taskHandler"
	"tasks-app-api/pkg/useCases/Helpers/responseHelper"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type TaskRouter struct {
	Handler taskHandler.Handler
}

func (tr *TaskRouter) GetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks, status := tr.Handler.GetAllTask()
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	w.WriteHeader(status.StatusCode())
	responseHelper.ResponseStatusChecker(w, resp)
}

func (tr *TaskRouter) GetTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(chi.URLParam(r, "taskId"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	task, status := tr.Handler.GetTask(taskId)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	w.WriteHeader(status.StatusCode())
	responseHelper.ResponseStatusChecker(w, resp)
}

func (tr *TaskRouter) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	status := tr.Handler.UpdateTask(task)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	w.WriteHeader(status.StatusCode())
	responseHelper.ResponseStatusChecker(w, resp)
}

func (tr *TaskRouter) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	status := tr.Handler.CreateTask(task)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseHelper.ResponseStatusChecker(w, INTERNAL_SERVER_ERROR)
		return
	}
	w.WriteHeader(status.StatusCode())
	responseHelper.ResponseStatusChecker(w, resp)
}

func (tr *TaskRouter) Routes() http.Handler {
	r := chi.NewRouter()

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"https://*", "http://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowOriginFunc:    func(r *http.Request, origin string) bool { return true },
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
		MaxAge:             300,
	}))

	r.Get("/{ taskId }", tr.GetTask)
	r.Get("/", tr.GetAllTask)

	r.Post("/", tr.CreateTask)

	r.Put("/", tr.UpdateTask)

	return r
}
