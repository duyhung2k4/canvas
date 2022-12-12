package model

type Admin struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Pass  string
	Email string
}

type Teacher struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Pass  string
	Email string
}

type Student struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Pass  string
	Email string
}
