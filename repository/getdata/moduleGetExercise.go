package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetExcercisesOfCourse(courseId int) ([]model.Exercise, error) {

	db := database.MethodDB()

	var listExcercise []model.Exercise

	err := db.Debug().
		Model(model.Exercise{}).
		Where("course_id = ?", courseId).
		Find(&listExcercise).Error

	if err != nil {
		fmt.Println(err)
	}

	return listExcercise, err

}

func GetOneExerciseOfCourse(courseId int, exerciseId int) (model.Exercise, error) {

	db := database.MethodDB()

	var exercise model.Exercise

	err := db.Debug().
		Model(model.Exercise{}).
		Where("id = ? AND course_id = ?", exerciseId, courseId).
		First(&exercise).Error

	if err != nil {
		fmt.Println(err)
	}

	return exercise, err
}
