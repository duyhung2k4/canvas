package model

type Request struct {
	Data    interface{}
	Message interface{}
	Err     error
	Success bool
}

type DataRequest struct {
	Text                string
	Token               string
	AccountTeacherId    Teacher
	AccountStudentId    Student
	SubjectId           Subject
	CourseId            Course
	FileId              File
	ExerciseId          Exercise
	CourseOfStudentId   CourseOfStudent
	AssignmentId        Assignment
	ChapterId           Chapter
	ListAccountTeacher  []Teacher
	ListAccountStudent  []Student
	ListSubject         []Subject
	ListCourse          []Course
	ListFile            []File
	ListExercise        []Exercise
	ListCourseOfStudent []CourseOfStudent
	ListAssignment      []Assignment
	ListChapter         []Chapter
}
