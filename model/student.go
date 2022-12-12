package model

type CourseOfStudent struct {
	Id        int `gorm:"primaryKey"`
	StudentId int
	CourseId  int
	Student   Student `gorm:"foreignKey:StudentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Course    Course  `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type AttendanceStatus struct {
	Id        int `gorm:"primaryKey"`
	StudentId int
	CourseId  int
	Absent    int
	Eligible  bool
	Student   Student `gorm:"foreignKey:StudentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Course    Course  `gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type Assignment struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	Link       string
	StudentId  int
	ExerciseId int
	Student    Student  `gorm:"foreignKey:StudentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Exercise   Exercise `gorm:"foreignKey:ExerciseId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
