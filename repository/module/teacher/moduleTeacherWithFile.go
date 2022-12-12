package teacher

import (
	datatbase "app/config"
	"app/model"
	"fmt"
)

func TeacherCreateFileForCourse(file model.File) error {

	db := datatbase.MethodDB()

	err := db.Debug().
		Model(model.File{}).
		Create(&file).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func TeacherUpdateFileForCourse(file model.File) error {

	db := datatbase.MethodDB()

	err := db.Debug().
		Model(model.File{}).
		Where("id = ? AND course_id = ?", file.Id, file.CourseId).
		Updates(&file).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func TeacherDeleteFileInCourse(file model.File) error {

	db := datatbase.MethodDB()

	err := db.Debug().
		Model(model.File{}).
		Where("id = ? AND course_id = ?", file.Id, file.CourseId).
		Delete(&file).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
