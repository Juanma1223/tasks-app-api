package responseHelper

import (
	"encoding/json"
	"net/http"
	"tasks-app-api/pkg/domain/response"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	response := response.Response{
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil
}

func ResponseStatusChecker(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		return
	}
}
