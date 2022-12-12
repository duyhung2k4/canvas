package checkexist

import (
	database "app/config"
	"app/model"
)

func CheckexistAssignmentOfExercise(assignment model.Assignment) error {

	db := database.MethodDB()

	var assignmentOfExercise model.Assignment

	err := db.Debug().
		Model(model.Assignment{}).
		Where("student_id = ? AND exercise_id = ?", assignment.StudentId, assignment.ExerciseId).
		First(&assignmentOfExercise).Error

	return err
}
