package student

import (
	database "app/config"
	"app/model"
	"fmt"
)

func StudentJoinCourse(courseOfStudent model.CourseOfStudent) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.CourseOfStudent{}).
		Create(&courseOfStudent).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
