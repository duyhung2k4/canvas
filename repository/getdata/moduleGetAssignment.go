package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAllAssignmentOfExercise(exerciseId int) ([]model.Assignment, error) {

	db := database.MethodDB()

	var assignments []model.Assignment

	err := db.Debug().
		Model(model.Assignment{}).
		Where("exercise_id = ?", exerciseId).
		Find(&assignments).Error

	if err != nil {
		fmt.Println(err)
	}

	return assignments, err
}

func GetAssignmentOfExercise(studentId int, exerciseId int) (model.Assignment, error) {

	db := database.MethodDB()

	var assignment model.Assignment

	err := db.Debug().
		Model(model.Assignment{}).
		Where("student_id = ? AND exercise_id = ?", studentId, exerciseId).
		First(&assignment).Error

	if err != nil {
		fmt.Println(err)
	}

	return assignment, err
}

func GetAssignmentOfExerciseForTeacher(exerciseId int, assignmentId int) (model.Assignment, error) {

	db := database.MethodDB()

	var assignment model.Assignment

	err := db.Debug().
		Model(model.Assignment{}).
		Where("id = ? AND exercise_id = ?", assignmentId, exerciseId).
		First(&assignment).Error

	if err != nil {
		fmt.Println(err)
	}

	return assignment, err
}
