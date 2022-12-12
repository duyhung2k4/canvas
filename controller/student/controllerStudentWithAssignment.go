package student

import (
	"app/errrequest"
	"app/model"
	"app/repository/checkexist"
	"app/repository/getdata"
	moduleStudent "app/repository/module/student"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// @Summary
// @Description Student create assignment
// @Tags Student with assignment
// @Accept json
// @Produce json
// @Param exercise_id path string true "Exercise id"
// @Router /student/my_course/{id}/exercise/{exercise_id}/assignment_create [post]
func StudentCreateAssignmentForExercise(w http.ResponseWriter, r *http.Request) {

	var assignment model.Assignment
	json.NewDecoder(r.Body).Decode(&assignment)

	_, claims, err1 := jwtauth.FromContext(r.Context())
	exerciseId, err2 := strconv.Atoi(chi.URLParam(r, "exercise_id"))

	if err1 != nil {
		errrequest.ErrToken(w, r, err1)
	} else if err2 != nil {
		errrequest.BadRequest(w, r, err2)
	} else {

		studentId := int(claims["id"].(float64))
		assignment.ExerciseId = exerciseId
		assignment.StudentId = studentId

		statusAssignment := checkexist.CheckexistAssignmentOfExercise(assignment)

		if statusAssignment == nil {
			errrequest.ArealyExist(w, r, "Exercise exist")
		} else {
			err := moduleStudent.StudentCreateAssignmentForExercise(assignment)

			if err != nil {
				errrequest.BadRequest(w, r, err)
			} else {
				result, _ := json.Marshal(model.Request{
					Data:    nil,
					Message: "Created Successfully!",
					Err:     nil,
					Success: true,
				})

				w.WriteHeader(http.StatusOK)
				w.Write(result)
			}
		}
	}
}

// @Summary
// @Description Student get assignment of exercise
// @Tags Student with assignment
// @Accept json
// @Produce json
// @Param exercise_id path string true "Exercise id"
// @Router /student/my_course/{id}/exercise/{exercise_id}/assignment [get]
func StudentGetAssignmentOfCourse(w http.ResponseWriter, r *http.Request) {

	exerciseId, err1 := strconv.Atoi(chi.URLParam(r, "exercise_id"))
	_, claims, err2 := jwtauth.FromContext(r.Context())

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.ErrToken(w, r, err2)
	} else {
		studentId := int(claims["id"].(float64))
		assignment, err := getdata.GetAssignmentOfExercise(studentId, exerciseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					AssignmentId: assignment,
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
// @Description Student update assignment
// @Tags Student with assignment
// @Accept json
// @Produce json
// @Param exercise_id path string true "Exercise Id"
// @Param req body model.Assignment false "Assignment need update" Format(model.Assignment)
// @Router /student/my_course/{id}/exercise/{exercise_id}/assignment_update [put]
func StudentUpdateAssignment(w http.ResponseWriter, r *http.Request) {

	var assignment model.Assignment
	json.NewDecoder(r.Body).Decode(&assignment)

	exerciseId, err1 := strconv.Atoi(chi.URLParam(r, "exercise_id"))
	_, claims, err2 := jwtauth.FromContext(r.Context())

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.ErrToken(w, r, err2)
	} else {
		studentId := int(claims["id"].(float64))
		assignment.StudentId = studentId
		assignment.ExerciseId = exerciseId

		err := moduleStudent.StudentUpdateAssignment(assignment)

		if err != nil {
			errrequest.BadRequest(w, r, err)
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
