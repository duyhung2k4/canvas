package admin

import (
	errRequest "app/errrequest"
	"app/model"
	"app/repository/checkexist"
	"app/repository/getdata"
	moduleAdmin "app/repository/module/admin"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Admin create subject
// @Tags Admin with subject
// @Accept json
// @Produce json
// @Param req body model.Subject false "Subject" Format(model.Subject)
// @Router /admin/subject/create [post]
func AdminCreateSubject(w http.ResponseWriter, r *http.Request) {

	var subject model.Subject

	json.NewDecoder(r.Body).Decode(&subject)

	statusSubject := checkexist.CheckexistSubject(subject)

	if statusSubject == nil {
		errRequest.ArealyExist(w, r, "Subject exist")
	} else {

		err := moduleAdmin.AdminCreateSubject(subject)

		if err != nil {
			errRequest.NotSuccess(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Creates Successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Admin get all subject
// @Tags Admin with subject
// @Accept json
// @Produce json
// @Router /admin/subject/list [get]
func AdminGetSubjects(w http.ResponseWriter, r *http.Request) {

	listSubject, err := getdata.GetAllSubject()

	if err != nil {
		errRequest.NotFoundRequest(w, r, err, "Not found list subject")
	} else {
		result, _ := json.Marshal(model.Request{
			Data: model.DataRequest{
				ListSubject: listSubject,
			},
			Message: "",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

// @Sumary
// @Description Admin get subject by id
// @Tags Admin with subject
// @Accept json
// @Produce json
// @Param id path string true "Subject id"
// @Router /admin/subject/list/{id} [get]
func AdminGetSubjectById(w http.ResponseWriter, r *http.Request) {

	paramIdSubject := chi.URLParam(r, "id")
	subjectId, err := strconv.Atoi(paramIdSubject)

	if err != nil {
		errRequest.BadRequest(w, r, err)
	} else {

		subject, err := getdata.GetSubjectById(subjectId)

		if err != nil {
			errRequest.NotFoundRequest(w, r, err, "subject not found")
		} else {

			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					SubjectId: subject,
				},
				Message: "",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Desrciption Admin update subject
// @Tags Admin with subject
// @Accept json
// @Produce json
// @Param req body model.Subject false "Subject need update" Format(model.Subject)
// @Router /admin/subject/update [put]
func AdminUpdateSubject(w http.ResponseWriter, r *http.Request) {

	var subject model.Subject

	json.NewDecoder(r.Body).Decode(&subject)

	statusSubject := checkexist.CheckexistSubject(subject)

	if statusSubject == nil {
		errRequest.ArealyExist(w, r, "Subject exist")
	} else {

		err := moduleAdmin.AdminUpdateSubject(subject)

		if err != nil {
			errRequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Updated Successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Admin delete subject
// @Tags Admin with subject
// @Accept json
// @Produce json
// @Param req body model.Subject false "Subject need delete" Format(mode.Subject)
// @Router /admin/subject/delete [delete]
func AdminDeleteSubject(w http.ResponseWriter, r *http.Request) {

	var subject model.Subject

	json.NewDecoder(r.Body).Decode(&subject)

	err := moduleAdmin.AdminDeleteSubject(subject)

	if err != nil {
		errRequest.BadRequest(w, r, err)
	} else {

		result, _ := json.Marshal(model.Request{
			Data:    nil,
			Message: "Deleted Successfully!",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
