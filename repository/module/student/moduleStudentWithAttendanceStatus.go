package student

import (
	database "app/config"
	"app/model"
	"fmt"
)

func CreateAttendanceStatusForStudent(attendanceStatus model.AttendanceStatus) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.AttendanceStatus{}).
		Create(&attendanceStatus).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
