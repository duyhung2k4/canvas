package admin

import (
	database "app/config"
	"app/model"
	"fmt"
)

func AdminCreateAccountTeacher(teacher model.Teacher) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Teacher{}).
		Create(&teacher).Error

	if err != nil {
		fmt.Println(err)
	}

	return err

}

func AdminUpdateAccountTeacher(accountTeacher model.Teacher) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Teacher{}).
		Where("id = ?", accountTeacher.Id).
		Updates(&accountTeacher).Error

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func AdminDeleteAccountTeacher(accountTeacher model.Teacher) error {

	db := database.MethodDB()

	err := db.Debug().
		Model(model.Teacher{}).
		Where("id = ?", accountTeacher.Id).
		Delete(&accountTeacher).Error

	if err != nil {

		fmt.Println(err)
	}

	return err

}
