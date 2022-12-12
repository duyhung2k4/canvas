package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAllAccountTeacher() ([]model.Teacher, error) {

	db := database.MethodDB()

	var listAccountTeacher []model.Teacher

	err := db.Debug().
		Model(model.Teacher{}).
		Find(&listAccountTeacher).Error

	if err != nil {

		fmt.Println(err)
	}

	return listAccountTeacher, err

}

func GetAccountTeacherById(teacherId int) (model.Teacher, error) {

	db := database.MethodDB()

	var accountteaCher model.Teacher

	err := db.Debug().
		Model(model.Teacher{}).
		Where("id = ?", teacherId).
		First(&accountteaCher).Error

	if err != nil {
		fmt.Println(err)
	}

	return accountteaCher, err
}

func GetAllAccountStudent() ([]model.Student, error) {

	db := database.MethodDB()

	var listAccountStudent []model.Student

	err := db.Debug().
		Model(model.Student{}).
		Find(&listAccountStudent).Error

	if err != nil {

		fmt.Println(err)
	}

	return listAccountStudent, err
}

func GetAccountStudentById(studentId int) (model.Student, error) {

	db := database.MethodDB()

	var student model.Student

	err := db.Debug().
		Model(model.Student{}).
		Where("id = ?", studentId).
		First(&student).Error

	if err != nil {
		fmt.Println(err)
	}

	return student, err
}
