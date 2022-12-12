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
// @Description Teacher create chapter for course
// @Tags Teacher with chapter
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param req body model.Chapter false "Chapter" Format(model.Chapter)
// @Router /teacher/courses/{id}/chapter/create [post]
func TeacherCreateChapter(w http.ResponseWriter, r *http.Request) {

	courseId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		var newChapter model.Chapter
		json.NewDecoder(r.Body).Decode(&newChapter)
		newChapter.CourseId = courseId

		err := moduleTeacher.TeacherCreateChapter(newChapter)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Created successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Teacher get all chapter of course
// @Tags Teacher with chapter
// @Accept json
// @Produce json
// @Param id path string true "Course Id"
// @Router /teacher/courses/{id}/chapter [get]
func TeacherGetAllChapterInCourse(w http.ResponseWriter, r *http.Request) {

	courseId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		listChapter, err := getdata.GetAllChapterOfCourse(courseId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ListChapter: listChapter,
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
// @Description Teacher get chapter id
// @Tags Teacher with chapter
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param chapter_id path string true "Chapter id"
// @Router /teacher/courses/{id}/chapter/{chapter_id} [get]
func TeacherGetChapterIdInCourse(w http.ResponseWriter, r *http.Request) {

	courseId, err1 := strconv.Atoi(chi.URLParam(r, "id"))
	chapterId, err2 := strconv.Atoi(chi.URLParam(r, "chapter_id"))

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else if err2 != nil {
		errrequest.BadRequest(w, r, err2)
	} else {

		chapter, err := getdata.GetChapterIdOfCourse(courseId, chapterId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					ChapterId: chapter,
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
// @Description Teacher update chapter
// @Tags Teacher with chapter
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param chapter_id path string true "Chapter id"
// @Router /teacher/courses/{id}/chapter/update [put]
func TeacherUpdateChapter(w http.ResponseWriter, r *http.Request) {

	courseId, err1 := strconv.Atoi(chi.URLParam(r, "id"))

	if err1 != nil {
		errrequest.BadRequest(w, r, err1)
	} else {

		var chapter model.Chapter
		json.NewDecoder(r.Body).Decode(&chapter)
		chapter.CourseId = courseId

		err := moduleTeacher.TeacherUpdateChapter(chapter)

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
// @Description Teacher delete chapter
// @Tags Teacher with chapter
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Router /teacher/courses/{id}/chapter/delete [delete]
func TeacherDeleteChapter(w http.ResponseWriter, r *http.Request) {

	courseId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {

		var chapter model.Chapter
		json.NewDecoder(r.Body).Decode(&chapter)
		chapter.CourseId = courseId

		err := moduleTeacher.TeacherDeleteChapter(chapter)

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
