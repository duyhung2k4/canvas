package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAllChapterOfCourse(courseId int) ([]model.Chapter, error) {

	db := database.MethodDB()

	var listChapter []model.Chapter

	err := db.Debug().
		Model(model.Chapter{}).
		Where("course_id = ?", courseId).
		Find(&listChapter).Error

	if err != nil {
		fmt.Println(err)
	}

	return listChapter, err
}

func GetChapterIdOfCourse(courseId int, chapterId int) (model.Chapter, error) {

	db := database.MethodDB()

	var chapter model.Chapter

	err := db.Debug().
		Model(model.Chapter{}).
		Where("id = ? AND course_id = ?", chapterId, courseId).
		First(&chapter).Error

	if err != nil {
		fmt.Println(err)
	}

	return chapter, err
}
