package checkexist

import (
	databse "app/config"
	"app/model"
)

func CheckexistCourse(course model.Course) error {

	db := databse.MethodDB()

	var courseInDB model.Course

	err := db.Debug().
		Model(model.Course{}).
		Where("code_course = ?", course.CodeCourse).
		First(&courseInDB).Error

	return err

}

func CheckexistCourseOfTeacher(teacherId int, courseId int) error {

	db := databse.MethodDB()

	var course model.Course

	err := db.Debug().
		Model(model.Course{}).
		Where("id = ? AND teacher_id = ?", courseId, teacherId).
		First(&course).Error

	return err
}
