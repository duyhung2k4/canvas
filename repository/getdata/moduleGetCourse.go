package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetListCourse() ([]model.Course, error) {

	db := database.MethodDB()

	var listCourse []model.Course

	err := db.Debug().
		Model(model.Course{}).
		Preload("Subject").
		Find(&listCourse).Error

	if err != nil {
		fmt.Println(err)
	}

	return listCourse, err
}

func GetCourseByTeacherId(teacherId int) ([]model.Course, error) {

	db := database.MethodDB()

	var courses []model.Course

	err := db.Debug().
		Model(model.Course{}).
		Where("teacher_id = ?", teacherId).
		Find(&courses).Error

	if err != nil {
		fmt.Println(err)
	}

	return courses, err
}

func GetCourseById(courseId int) (model.Course, error) {

	db := database.MethodDB()

	var courseByID model.Course

	err := db.Debug().
		Model(model.Course{}).
		Preload("Subject").
		Where("id = ?", courseId).
		First(&courseByID).Error

	if err != nil {
		fmt.Println(err)
	}

	return courseByID, err
}

func GetCourseByIdOfTeacher(teacherId int, courseId int) (model.Course, error) {

	db := database.MethodDB()

	var course model.Course

	err := db.Debug().
		Model(model.Course{}).
		Where("id = ? AND teacher_id = ?", courseId, teacherId).
		First(&course).Error

	if err != nil {
		fmt.Println(err)
	}

	return course, err
}
