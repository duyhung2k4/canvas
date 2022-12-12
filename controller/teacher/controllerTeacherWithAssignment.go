package teacher

import (
	"app/errrequest"
	"app/model"
	"app/repository/getdata"
	moduleTeacher "app/repository/module/teacher"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Teacher get assignment of exercise
// @Tags Teacher with assignment
// @Accept json
// @Produce json
// @Param exercise_id path string true "exercise id"
// @Router /teacher/courses/{id}/exercise/{exercise_id}/assignment [get]
func TeacherGetAllAssignmentOfExercise(w http.ResponseWriter, r *http.Request) {

	exerciseId, err := strconv.Atoi(chi.URLParam(r, "exercise_id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		listAssignment, err := getdata.GetAllAssignmentOfExercise(exerciseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ListAssignment: listAssignment,
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
// @Description Teacher get assignment
// @Tags Teacher with assignment
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param exercise_id path string true "Exercise id"
// @Param assignment_id path string true "Assignment id"
// @Router /teacher/courses/{id}/exercise/{exercise_id}/assignment/{assignment_id} [get]
func TeacherGetAssignmentIdInExercise(w http.ResponseWriter, r *http.Request) {

	exerciseId, err1 := strconv.Atoi(chi.URLParam(r, "exercise_id"))
	assignmentId, err2 := strconv.Atoi(chi.URLParam(r, "assignment_id"))

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.BadRequest(w, r, err2)
	} else {

		assignment, err := getdata.GetAssignmentOfExerciseForTeacher(exerciseId, assignmentId)

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
// @Description Teacher update assignment
// @Tags Teacher with assignment
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param exercise_id path string true "Exercise id"
// @Router /teacher/courses/{id}/exercise/{exercise_id}/assignment/update [put]
func TeacherUpdateAssignment(w http.ResponseWriter, r *http.Request) {

	var assignment model.Assignment
	json.NewDecoder(r.Body).Decode(&assignment)

	exerciseId, err := strconv.Atoi(chi.URLParam(r, "exercise_id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		assignment.ExerciseId = exerciseId
		err := moduleTeacher.TeacherUpdateAssignment(assignment)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Updated successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}
