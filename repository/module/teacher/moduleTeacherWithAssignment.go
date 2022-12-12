package teacher

import (
	database "app/config"
	"app/model"
	"fmt"
)

func TeacherUpdateAssignment(assignment model.Assignment) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Assignment{}).
		Where("id = ?", assignment.Id).
		Updates(&assignment).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
