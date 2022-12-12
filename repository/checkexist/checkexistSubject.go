package checkexist

import (
	database "app/config"
	"app/model"
	"fmt"
)

func CheckexistSubject(subject model.Subject) error {

	db := database.MethodDB()

	var subjectInDB model.Subject

	err := db.Debug().
		Model(model.Subject{}).
		Where("code_subject = ?", subject.CodeSubject).
		First(&subjectInDB).Error

	if err != nil {
		fmt.Println(err)
	}

	return err

}

func CheckexistSubjectById(subjectId int) error {

	db := database.MethodDB()

	var subject model.Subject

	err := db.Debug().
		Model(model.Subject{}).
		Where("id = ?", subjectId).
		First(&subject).Error

	return err
}
