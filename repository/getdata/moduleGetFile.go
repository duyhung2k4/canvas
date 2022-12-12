package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAllFileOfCourse(courseId int) ([]model.File, error) {

	db := database.MethodDB()

	var listFile []model.File

	err := db.Debug().
		Model(model.File{}).
		Where("course_id = ?", courseId).
		Find(&listFile).Error

	if err != nil {
		fmt.Println(err)
	}

	return listFile, err
}

func GetFileOfCourse(courseId int, fileId int) (model.File, error) {

	db := database.MethodDB()

	var file model.File

	err := db.Debug().
		Model(model.File{}).
		Where("id = ? AND course_id = ?", fileId, courseId).
		First(&file).Error

	if err != nil {
		fmt.Println(err)
	}

	return file, err
}
