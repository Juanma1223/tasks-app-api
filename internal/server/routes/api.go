package routes

import (
	"net/http"
	templaterepository "template/internal/data/infrastructure/templateRepository"
	templatehandler "template/pkg/useCases/Handlers/templateHandler"

	"github.com/go-chi/chi"
)

var (
	INTERNAL_SERVER_ERROR = []byte("500: Internal Server Error")
	ERR_ALREADY_COMMITTED = "already been committed"
)

func New() http.Handler {
	r := chi.NewRouter()

	tr := TemplateRouter{
		Handler: templatehandler.TemplateHandler{
			Repository: templaterepository.TemplateRepository{},
		},
	}

	r.Mount("/template", tr.Routes())
	return r
}
