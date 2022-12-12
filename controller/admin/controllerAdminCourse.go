package admin

import (
	"app/errrequest"
	"app/model"
	"app/repository/checkexist"
	"app/repository/getdata"
	"encoding/json"
	"net/http"
	"strconv"

	moduleAdmin "app/repository/module/admin"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Admin create course
// @Tags Admin with course
// @Accept json
// @Produce json
// @Param req body model.Course false "Course" Format(model.Course)
// @Router /admin/course/create [post]
func AdminCreateCourse(w http.ResponseWriter, r *http.Request) {

	var course model.Course

	json.NewDecoder(r.Body).Decode(&course)

	err := checkexist.CheckexistCourse(course)
	statusTeacher := checkexist.CheckexistAccountTeacherById(course.TeacherId)
	statusSubject := checkexist.CheckexistSubjectById(course.SubjectId)

	if err == nil {
		errrequest.NotFoundRequest(w, r, err, "Course exist")
	} else if statusTeacher != nil {
		errrequest.NotFoundRequest(w, r, statusTeacher, "Not exist teacher")
	} else if statusSubject != nil {
		errrequest.NotFoundRequest(w, r, statusSubject, "Not exist subject")
	} else {

		err := moduleAdmin.AdminCreateCourse(course)

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
// @Description Admin get list course
// @Tags Admin with course
// @Accept json
// @Produce json
// @Router /admin/course/list [get]
func AdminGetCourses(w http.ResponseWriter, r *http.Request) {

	listCourse, err := getdata.GetListCourse()

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		result, _ := json.Marshal(model.Request{
			Data: model.DataRequest{
				ListCourse: listCourse,
			},
			Message: "",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

// @Summary
// @Description Admin get course by id
// @Tags Admin with course
// @Accept json
// @Produce json
// @Param id path string true "Course Id"
// @Router /admin/course/list/{id} [get]
func AdminGetCourseById(w http.ResponseWriter, r *http.Request) {

	paramCourseId := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseId)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		course, err := getdata.GetCourseById(courseId)

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

// @Sumarry
// @Description Admin update course
// @Tags Admin with course
// @Accept json
// @Produce json
// @Param req body model.Course false "Course need update" Format(model.Course)
// @Router /admin/course/update [put]
func AdminUpdateCourse(w http.ResponseWriter, r *http.Request) {

	var course model.Course

	json.NewDecoder(r.Body).Decode(&course)

	statusCourse := checkexist.CheckexistCourse(course)
	statusTeacher := checkexist.CheckexistAccountTeacherById(course.TeacherId)
	statusSubject := checkexist.CheckexistSubjectById(course.SubjectId)

	if statusCourse == nil {
		errrequest.ArealyExist(w, r, "Exist course")
	} else if statusTeacher != nil {
		errrequest.NotFoundRequest(w, r, statusCourse, "Not exist teacher")
	} else if statusSubject != nil {
		errrequest.NotFoundRequest(w, r, statusSubject, "Not exist subject")
	} else {

		err := moduleAdmin.AdminUpdateCourse(course)

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

// @Summary
// @Description Admin delete course
// @Tags Admin with course
// @Accept json
// @Produce json
// @Param req body model.Course false "Course need delete" Format(model.Course)
// @Router /admin/course/delete [delete]
func AdminDeleteCourse(w http.ResponseWriter, r *http.Request) {

	var course model.Course

	json.NewDecoder(r.Body).Decode(&course)

	err := moduleAdmin.AdminDeleteCourse(course)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		result, _ := json.Marshal(model.Request{
			Data:    nil,
			Message: "Deleted successully!",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
