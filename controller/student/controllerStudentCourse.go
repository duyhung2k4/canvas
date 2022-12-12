package student

import (
	"app/errrequest"
	"app/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"app/repository/getdata"
	moduleStudent "app/repository/module/student"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// @Summary
// @Description Student join course
// @Tags Student with course
// @Accept json
// @Produce json
// @Param req body model.CourseOfStudent false "Course of student" Format(model.CourseOfStudent)
// @Router /student/joincourse [post]
func StudentJoinCourse(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		errrequest.ErrToken(w, r, err)
	} else {

		studentId := int(claims["id"].(float64))

		var courseOfStudent model.CourseOfStudent
		var attendanceStatus model.AttendanceStatus

		json.NewDecoder(r.Body).Decode(&courseOfStudent)

		fmt.Println(courseOfStudent)

		courseOfStudent.StudentId = studentId
		attendanceStatus.CourseId = courseOfStudent.CourseId
		attendanceStatus.StudentId = studentId
		attendanceStatus.Absent = 0
		attendanceStatus.Eligible = true

		err := moduleStudent.StudentJoinCourse(courseOfStudent)
		err2 := moduleStudent.CreateAttendanceStatusForStudent(attendanceStatus)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else if err2 != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Joined Successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Student get all course
// @Tags Student with course
// @Accept json
// @Produce json
// @Router /student/my_course [get]
func StudentGetAllCourse(w http.ResponseWriter, r *http.Request) {

	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		errrequest.ErrToken(w, r, err)
	} else {

		studentId := int(claims["id"].(float64))

		listCourseOfStudent, err := getdata.GetAllCourseOfStudent(studentId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ListCourseOfStudent: listCourseOfStudent,
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
// @Description Student get course of student by id
// @Tags Student with course
// @Accept json
// @Produce json
// @Param id path string true "Course of student id"
// @Router /student/my_course/{id} [get]
func StudentGetCourseById(w http.ResponseWriter, r *http.Request) {

	_, claims, err1 := jwtauth.FromContext(r.Context())
	courseId, err2 := strconv.Atoi(chi.URLParam(r, "id"))

	if err1 != nil {
		errrequest.ErrToken(w, r, err1)
	} else if err2 != nil {
		errrequest.BadRequest(w, r, err2)
	} else {
		studentId := int(claims["id"].(float64))

		courseOfStudent, err := getdata.GetCourseOfStudentById(studentId, courseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					CourseOfStudentId: courseOfStudent,
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
