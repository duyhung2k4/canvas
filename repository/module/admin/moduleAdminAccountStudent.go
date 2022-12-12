package admin

import (
	database "app/config"
	"app/model"
	"fmt"
)

func AdminCreateAccountStudent(student model.Student) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Student{}).
		Create(&student).Error

	return err
}

func AdminUpdateAccountStudent(student model.Student) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Student{}).
		Where("id = ?", student.Id).
		Updates(&student).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func AdminDeleteAccountStudent(student model.Student) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Student{}).
		Where("id = ?", student.Id).
		Delete(&student).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
