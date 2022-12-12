package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAllCourseOfStudent(studentId int) ([]model.CourseOfStudent, error) {

	db := database.MethodDB()

	var listCourseOfStudent []model.CourseOfStudent

	err := db.Debug().
		Model(model.CourseOfStudent{}).
		Where("student_id = ?", studentId).
		Preload("Course").
		Find(&listCourseOfStudent).Error

	if err != nil {
		fmt.Println(err)
	}

	return listCourseOfStudent, err
}

func GetCourseOfStudentById(studentId int, courseOfStudentId int) (model.CourseOfStudent, error) {

	db := database.MethodDB()

	var courseOfStudent model.CourseOfStudent

	err := db.Debug().
		Model(model.CourseOfStudent{}).
		Where("id = ? AND student_id = ?", courseOfStudentId, studentId).
		Preload("Course").
		First(&courseOfStudent).Error

	if err != nil {
		fmt.Println(err)
	}

	return courseOfStudent, err
}
