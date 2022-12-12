package admin

import (
	database "app/config"
	"app/model"
	"fmt"
)

func AdminCreateCourse(course model.Course) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Course{}).
		Create(&course).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func AdminUpdateCourse(course model.Course) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Course{}).
		Where("id = ?", course.Id).
		Updates(&course).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func AdminDeleteCourse(course model.Course) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Course{}).
		Where("id = ?", course.Id).
		Delete(&course).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
