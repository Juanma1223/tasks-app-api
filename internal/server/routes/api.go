package routes

import (
	"net/http"
	"tasks-app-api/internal/data/infrastructure/taskRepository"
	"tasks-app-api/pkg/useCases/Handlers/taskHandler"

	"github.com/go-chi/chi"
)

var (
	INTERNAL_SERVER_ERROR = []byte("500: Internal Server Error")
	ERR_ALREADY_COMMITTED = "already been committed"
)

func New() http.Handler {
	r := chi.NewRouter()

	tr := TaskRouter{
		Handler: &taskHandler.TaskHandler{
			Repository: &taskRepository.TaskRepository{},
		},
	}

	r.Mount("/tasks", tr.Routes())
	return r
}
