package teacher

import (
	database "app/config"
	"app/model"
	"fmt"
)

func TeacherCreateChapter(chapter model.Chapter) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Chapter{}).
		Create(&chapter).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func TeacherUpdateChapter(chapter model.Chapter) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Chapter{}).
		Where("id = ? AND course_id = ?", chapter.Id, chapter.CourseId).
		Updates(&chapter).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func TeacherDeleteChapter(chapter model.Chapter) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Chapter{}).
		Where("id = ? AND course_id = ?", chapter.Id, chapter.CourseId).
		Delete(&chapter).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
