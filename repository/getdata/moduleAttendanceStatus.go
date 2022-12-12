package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAttendanceStatus(attendance_id int) (model.AttendanceStatus, error) {

	db := database.MethodDB()

	var attendanceStatus model.AttendanceStatus

	err := db.Debug().
		Model(model.AttendanceStatus{}).
		Where("id = ?", attendance_id).
		Preload("Course").
		First(&attendanceStatus).Error

	if err != nil {
		fmt.Println(err)
	}

	return attendanceStatus, err

}
