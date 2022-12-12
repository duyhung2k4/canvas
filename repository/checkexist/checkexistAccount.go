package checkexist

import (
	database "app/config"
	"app/model"
	"fmt"
)

func CheckexistAccountAdmin(admin model.Admin) (model.Admin, error) {

	db := database.MethodDB()

	var adminInDB model.Admin

	err := db.Debug().
		Model(model.Admin{}).
		Where("email = ? AND pass = ?", admin.Email, admin.Pass).
		First(&adminInDB).Error

	if err != nil {
		fmt.Println(err)
	}

	return adminInDB, err
}

func CheckexistAccounTeacher(teacher model.Teacher) (model.Teacher, error) {

	db := database.MethodDB()

	var accountTeacher model.Teacher

	err := db.Debug().
		Model(model.Teacher{}).
		Where("email = ? AND pass = ?", teacher.Email, teacher.Pass).
		First(&accountTeacher).Error

	if err != nil {
		fmt.Println(err)
	}

	return accountTeacher, err
}

func CheckexistAccountStudent(student model.Student) (model.Student, error) {

	db := database.MethodDB()

	var accountStudent model.Student

	err := db.Debug().
		Model(model.Student{}).
		Where("email = ? AND pass = ?", student.Email, student.Pass).
		First(&accountStudent).Error

	if err != nil {
		fmt.Println(err)
	}

	return accountStudent, err
}

func CheckexistAccountTeacherById(teacherId int) error {

	db := database.MethodDB()

	var teacher model.Teacher

	err := db.Debug().
		Model(model.Teacher{}).
		Where("id = ?", teacherId).
		First(&teacher).Error

	return err
}
