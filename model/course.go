package model

type Subject struct {
	Id          int `gorm:"primaryKey"`
	Name        string
	CodeSubject string
}

type Course struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	CodeCourse string
	Lessons    int
	TeacherId  int
	SubjectId  int
	Teacher    Teacher `gorm:"foreignKey:TeacherId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Subject    Subject `gorm:"foreignKey:SubjectId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type File struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Link     string
	CourseId int
	Course   Course `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Exercise struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Link     string
	CourseId int
	Course   Course `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Chapter struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Link     string
	CourseId int
	Course   Course `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
