package templatehandler

import (
	templateRepository "template/internal/data/infrastructure/templateRepository"
)

type TemplateHandler struct {
	Repository templateRepository.Repository
}

type Handler interface {
}
