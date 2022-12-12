package teacher

import (
	database "app/config"
	"app/model"
	"fmt"
)

func TeacherCreateExercise(exercise model.Exercise) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Exercise{}).
		Where("course_id = ?", exercise.CourseId).
		Create(&exercise).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func TeacherUpdateExercise(exercise model.Exercise) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Exercise{}).
		Where("id = ? AND course_id = ?", exercise.Id, exercise.CourseId).
		Updates(&exercise).Error

	if err != nil {
		fmt.Println(err)
	}

	return err

}

func TeacherDeleteExercise(exercise model.Exercise) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Exercise{}).
		Where("id = ? AND course_id = ?", exercise.Id, exercise.CourseId).
		Delete(&exercise).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
