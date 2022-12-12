package student

import (
	database "app/config"
	"app/model"
	"fmt"
)

func StudentCreateAssignmentForExercise(assignment model.Assignment) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Assignment{}).
		Create(&assignment).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func StudentUpdateAssignment(assignment model.Assignment) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Assignment{}).
		Where("student_id = ? AND exercise_id = ?", assignment.StudentId, assignment.ExerciseId).
		Updates(&assignment).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
