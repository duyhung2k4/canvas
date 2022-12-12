package student

import (
	"app/errrequest"
	"app/model"
	"app/repository/getdata"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// @Summary
// @Description Student get exercise of course
// @Tags Student with exercise
// @Accept json
// @Produce json
// @Param id path string true "Course Of Student id"
// @Router /student/my_course/{id}/exercise [get]
func StudentAllGetExerciseofCourse(w http.ResponseWriter, r *http.Request) {

	courseOfStudentId, err1 := strconv.Atoi(chi.URLParam(r, "id"))
	_, claims, err2 := jwtauth.FromContext(r.Context())

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.ErrToken(w, r, err2)
	} else {
		studentId := int(claims["id"].(float64))
		courseOfStudentById, err := getdata.GetCourseOfStudentById(studentId, courseOfStudentId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			listExerciseofCourse, err := getdata.GetExcercisesOfCourse(courseOfStudentById.CourseId)

			if err != nil {
				errrequest.BadRequest(w, r, err)
			} else {
				result, _ := json.Marshal(model.Request{
					Data: model.DataRequest{
						ListExercise: listExerciseofCourse,
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

}

// @Summary
// @Description Student get exercise of course by id
// @Tags Student with exercise
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param exercise_id path string true "Exercise Id"
// @Router /student/my_course/{id}/exercise/{exercise_id} [get]
func StudentGetOneExerciseOfCourse(w http.ResponseWriter, r *http.Request) {

	courseOfStudentId, err1 := strconv.Atoi(chi.URLParam(r, "id"))
	_, claims, err2 := jwtauth.FromContext(r.Context())
	exerciseId, err3 := strconv.Atoi(chi.URLParam(r, "exercise_id"))

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.ErrToken(w, r, err2)
	} else if err3 != nil {
		errrequest.BadRequest(w, r, err3)
	} else {

		studentId := int(claims["id"].(float64))
		courseOfStudentById, err := getdata.GetCourseOfStudentById(studentId, courseOfStudentId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {

			exercisebyId, err := getdata.GetOneExerciseOfCourse(courseOfStudentById.CourseId, exerciseId)

			if err != nil {
				errrequest.BadRequest(w, r, err)
			} else {

				result, _ := json.Marshal(model.Request{
					Data: model.DataRequest{
						ExerciseId: exercisebyId,
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
}
