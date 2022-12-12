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
// @Description Teacher create file for course
// @Tags Teacher with file
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param req body model.File false "File" Format(model.File)
// @Router /teacher/courses/{id}/file/create [post]
func TeacherCreateFileForCourse(w http.ResponseWriter, r *http.Request) {

	var file model.File
	json.NewDecoder(r.Body).Decode(&file)

	paramCourseid := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseid)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		file.CourseId = courseId
		err := moduleTeacher.TeacherCreateFileForCourse(file)

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
// @Description Teacher get all file of course
// @Tags Teacher with file
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Router /teacher/courses/{id}/file [get]
func TeacherGetFileOfCourse(w http.ResponseWriter, r *http.Request) {

	paramCourseid := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseid)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		listFile, err := getdata.GetAllFileOfCourse(courseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ListFile: listFile,
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
// @Description teacher get a file of course
// @Tags Teacher with file
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param file_id path string true "file id"
// @Router /teacher/courses/{id}/file/{file_id} [get]
func TeacherOneGetFileOfCourse(w http.ResponseWriter, r *http.Request) {

	paramCourseId := chi.URLParam(r, "id")
	paramFileId := chi.URLParam(r, "file_id")

	courseId, err1 := strconv.Atoi(paramCourseId)
	fileId, err2 := strconv.Atoi(paramFileId)

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.BadRequest(w, r, err2)
	} else {

		file, err := getdata.GetFileOfCourse(courseId, fileId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					FileId: file,
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
// @Description Teacher update file of course
// @Tags Teacher with file
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param req body model.File false "File need update" Format(model.File)
// @Router /teacher/courses/{id}/file/update [put]
func TeacherUpdateFile(w http.ResponseWriter, r *http.Request) {

	var file model.File
	json.NewDecoder(r.Body).Decode(&file)

	paramCourseId := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseId)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		file.CourseId = courseId

		err := moduleTeacher.TeacherUpdateFileForCourse(file)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Updated Successfully",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Teacher delete file of course
// @Tags Teacher with file
// @Accept json
// @Produce json
// @Param id path string true "course id"
// @Param req body model.File false "File need delete" Format(model.File)
// @Router /teacher/courses/{id}/file/delete [delete]
func TeacherDeleteFileOfCourse(w http.ResponseWriter, r *http.Request) {

	paramCourseId := chi.URLParam(r, "id")
	courseId, err := strconv.Atoi(paramCourseId)

	var file model.File
	json.NewDecoder(r.Body).Decode(&file)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		file.CourseId = courseId

		err := moduleTeacher.TeacherDeleteFileInCourse(file)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Deleted successfully",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}
