package checkexist

import (
	database "app/config"
	"app/model"
	"fmt"
)

func CheckexistEmailTeacher(teacher model.Teacher) error {

	db := database.MethodDB()

	var accountTeacher model.Teacher

	err := db.Debug().
		Model(model.Teacher{}).
		Where("email = ?", teacher.Email).
		First(&accountTeacher).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func CheckexistEmailStudent(student model.Student) error {

	db := database.MethodDB()

	var accountStudent model.Student

	err := db.Debug().
		Model(model.Student{}).
		Where("email = ?", student.Email).
		First(&accountStudent).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
