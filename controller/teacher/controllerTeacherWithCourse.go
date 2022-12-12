package teacher

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
// @Description Course of teacher
// @Tags Teacher with course
// @Accept json
// @produce json
// @Router /teacher/courses [get]
func TeacherGetListCourse(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		errrequest.ErrToken(w, r, err)
	} else {
		teacherId := int(claims["id"].(float64))
		listCourseOfTeacher, err := getdata.GetCourseByTeacherId(teacherId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ListCourse: listCourseOfTeacher,
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
// @Description Course of teacher by id
// @Tags Teacher with course
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Router /teacher/courses/{id} [get]
func TeacherGetCourseById(w http.ResponseWriter, r *http.Request) {

	var course model.Course
	json.NewDecoder(r.Body).Decode(&course)

	paramCourseId := chi.URLParam(r, "id")
	courseId, errCourseId := strconv.Atoi(paramCourseId)
	_, claims, err := jwtauth.FromContext(r.Context())

	if errCourseId != nil {
		errrequest.BadRequest(w, r, errCourseId)
	} else if err != nil {
		errrequest.ErrToken(w, r, err)
	} else {
		teacherId := int(claims["id"].(float64))

		course, err := getdata.GetCourseByIdOfTeacher(teacherId, courseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					CourseId: course,
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
