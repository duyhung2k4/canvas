package teacher

import (
	database "app/config"
	"app/model"
	"app/repository/getdata"
	"fmt"
)

func TeacherUpdateAttendanceStatus(attendance_id int) error {

	db := database.MethodDB()

	attendanceStatus, err := getdata.GetAttendanceStatus(attendance_id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	absent := attendanceStatus.Absent

	var err2 error

	if ((absent+1)*100)/(attendanceStatus.Course.Lessons) > 10 {
		err2 = db.Debug().
			Model(model.AttendanceStatus{}).
			Where("id = ?", attendanceStatus.Id).
			Update("eligible", false).Error
	} else {
		err2 = db.Debug().
			Model(model.AttendanceStatus{}).
			Where("id = ?", attendance_id).
			Update("absent", absent+1).Error
	}

	if err2 != nil {
		fmt.Println(err2)
	}

	return err2
}
