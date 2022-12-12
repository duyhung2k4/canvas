package getdata

import (
	database "app/config"
	"app/model"
	"fmt"
)

func GetAllSubject() ([]model.Subject, error) {

	db := database.MethodDB()

	var listSubject []model.Subject

	err := db.Debug().
		Model(model.Subject{}).
		Find(&listSubject).Error

	if err != nil {
		fmt.Println(err)
	}

	return listSubject, err
}

func GetSubjectById(subjectId int) (model.Subject, error) {

	db := database.MethodDB()

	var subject model.Subject

	err := db.Debug().
		Model(model.Subject{}).
		Where("id = ?", subjectId).
		First(&subject).Error

	if err != nil {
		fmt.Println(err)
	}

	return subject, err
}
