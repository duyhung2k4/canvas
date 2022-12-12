package middleware

import (
	"app/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

func badRequest(w http.ResponseWriter, r *http.Request, err error) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: err,
		Success: false,
	})

	render.JSON(w, r, result)
}

func notIsPermission(w http.ResponseWriter, r *http.Request, message string) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: message,
		Success: false,
	})

	render.JSON(w, r, result)
}
