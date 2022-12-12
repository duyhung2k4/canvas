package errrequest

import (
	"app/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
)

func ArealyExist(w http.ResponseWriter, r *http.Request, message string) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Err:     nil,
		Message: message,
		Success: false,
	})

	w.Write(result)
	render.JSON(w, r, result)
}

func NotFoundRequest(w http.ResponseWriter, r *http.Request, err error, message string) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: message,
		Err:     err,
		Success: false,
	})

	w.WriteHeader(http.StatusNotFound)
	render.JSON(w, r, result)
}

func ErrToken(w http.ResponseWriter, r *http.Request, err error) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: "",
		Err:     err,
		Success: false,
	})

	w.WriteHeader(http.StatusBadRequest)
	render.JSON(w, r, result)
}

func ExistEmailRequest(w http.ResponseWriter, r *http.Request, message string) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: message,
		Err:     nil,
		Success: false,
	})

	w.Write(result)
	render.JSON(w, r, result)
}

func NotSuccess(w http.ResponseWriter, r *http.Request, err error) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: "",
		Err:     err,
		Success: false,
	})

	w.Write(result)
	render.JSON(w, r, result)
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) {

	result, _ := json.Marshal(model.Request{
		Data:    nil,
		Message: "",
		Err:     err,
		Success: false,
	})

	w.WriteHeader(http.StatusBadRequest)
	w.Write(result)
	render.JSON(w, r, result)
}
