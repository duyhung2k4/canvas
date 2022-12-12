package teacher

import (
	"app/errrequest"
	"app/model"
	"encoding/json"
	"net/http"
	"strconv"

	moduleTeacher "app/repository/module/teacher"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Teacher update attendance status student
// @Tags Teacher with attendance status
// @Accept json
// @Produce json
// @Param id path string true "Course id"
// @Param attendance_id path string true "Attendance status id"
// @Router /teacher/courses/{id}/attendance/{attendance_id} [put]
func TeacherUpdateAttendanceStatus(w http.ResponseWriter, r *http.Request) {

	attendanceId, err := strconv.Atoi(chi.URLParam(r, "attendance_id"))

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		err := moduleTeacher.TeacherUpdateAttendanceStatus(attendanceId)
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
