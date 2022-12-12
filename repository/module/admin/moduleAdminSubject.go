package admin

import (
	database "app/config"
	"app/model"
	"fmt"
)

func AdminCreateSubject(subject model.Subject) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Subject{}).
		Create(&subject).Error

	if err != nil {
		fmt.Println(err)
	}

	return err

}

func AdminUpdateSubject(subject model.Subject) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Subject{}).
		Where("id = ?", subject.Id).
		Updates(&subject).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func AdminDeleteSubject(subject model.Subject) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Subject{}).
		Where("id = ?", subject.Id).
		Delete(&subject).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}
