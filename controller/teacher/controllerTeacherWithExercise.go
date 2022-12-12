package teacher

import (
	"app/errrequest"
	"app/model"
	"encoding/json"
	"net/http"
	"strconv"

	"app/repository/getdata"
	moduleTeacher "app/repository/module/teacher"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Teacher create exercise
// @Tags Teacher with exercise
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param req body model.Exercise false "Exercise" Format(model.Exercise)
// @Router /teacher/courses/{id}/exercise/create [post]
func TeacherCreateExercise(w http.ResponseWriter, r *http.Request) {

	paramCourseId := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseId)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		var exercise model.Exercise
		json.NewDecoder(r.Body).Decode(&exercise)

		exercise.CourseId = courseId

		err := moduleTeacher.TeacherCreateExercise(exercise)

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

// @Summary
// @Description Teacher get exercises of course
// @Tags Teacher with exercise
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Router /teacher/courses/{id}/exercise [get]
func TeacherGetAllExerciseOfCourse(w http.ResponseWriter, r *http.Request) {

	paramCourseId := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseId)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		listExercise, err := getdata.GetExcercisesOfCourse(courseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ListExercise: listExercise,
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
// @Description Teacher get exercise by id
// @Tags Teacher with exercise
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param exercise_id path string true "execie id"
// @Router /teacher/courses/{id}/exercise/{exercise_id} [get]
func TeacherGetExerciseById(w http.ResponseWriter, r *http.Request) {

	paramExerciseId := chi.URLParam(r, "exercise_id")
	paramCourseId := chi.URLParam(r, "id")

	exerciseid, err1 := strconv.Atoi(paramExerciseId)
	courseId, err2 := strconv.Atoi(paramCourseId)

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.BadRequest(w, r, err2)
	} else {

		exercise, err := getdata.GetOneExerciseOfCourse(courseId, exerciseid)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ExerciseId: exercise,
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

// @Summry
// @Description Teacher update exercise of course
// @Tags Teacher with exercise
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param req body model.Exercise false "exercise need update" Format(model.Exercise)
// @Router /teacher/courses/{id}/exercise/update [put]
func TeacherUpdateExercise(w http.ResponseWriter, r *http.Request) {

	paramCourseId := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseId)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		var exercise model.Exercise
		json.NewDecoder(r.Body).Decode(&exercise)

		exercise.CourseId = courseId

		err := moduleTeacher.TeacherUpdateExercise(exercise)

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

// @Summary
// @Description Teacher delete exercise
// @Tags Teacher with exercise
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param req body model.Exercise false "exercise need delete" Format(model.Exercise)
// @Router /teacher/courses/{id}/exercise/delete [delete]
func TeacherDeleteExercise(w http.ResponseWriter, r *http.Request) {

	courseId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		var exercise model.Exercise
		exercise.CourseId = courseId

		err := moduleTeacher.TeacherDeleteExercise(exercise)

		if err != nil {
			errrequest.BadRequest(w, r, err)
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
}
