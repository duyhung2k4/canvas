package router

import (
	controllerAdmin "app/controller/admin"
	controllerLogin "app/controller/login"
	controllerStudent "app/controller/student"
	controllerTeacher "app/controller/teacher"
	"app/middleware"
	"fmt"
	"net/http"
	"os"

	_ "app/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func App() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := chi.NewRouter()
	tokenAuth := jwtauth.New("HS256", []byte("token"), nil)
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	app.Use(cors.Handler)

	app.Route("/login", func(r chi.Router) {

		r.Post("/admin", controllerLogin.AdminLogin)
		r.Post("/teacher", controllerLogin.LoginTeacher)
		r.Post("/student", controllerLogin.LoginStudent)
	})

	app.Route("/admin", func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middleware.AuthenticationAdmin)

		r.Route("/teacher", func(r2 chi.Router) {

			r2.Get("/accounts", controllerAdmin.AdminGetAccountTeachers)
			r2.Get("/accounts/{id}", controllerAdmin.AdminGetAccountTeacherById)
			r2.Post("/create", controllerAdmin.AdminCreateAccountAdmin)
			r2.Put("/update", controllerAdmin.AdminUpdateAccountTeacher)
			r2.Delete("/delete", controllerAdmin.AdminDeleteAccountTeacher)
		})

		r.Route("/student", func(r2 chi.Router) {

			r2.Get("/accounts", controllerAdmin.AdminGetAccountStudents)
			r2.Get("/accounts/{id}", controllerAdmin.AdminGetAccountStudentById)
			r2.Post("/create", controllerAdmin.AdminCreateAccountStudent)
			r2.Put("/update", controllerAdmin.AdminUpdateAccountStudent)
			r2.Delete("/delete", controllerAdmin.AdminDeleteAccountStudent)
		})

		r.Route("/subject", func(r2 chi.Router) {

			r2.Get("/list", controllerAdmin.AdminGetSubjects)
			r2.Get("/list/{id}", controllerAdmin.AdminGetSubjectById)
			r2.Post("/create", controllerAdmin.AdminCreateSubject)
			r2.Put("/update", controllerAdmin.AdminUpdateSubject)
			r2.Delete("/delete", controllerAdmin.AdminDeleteSubject)
		})

		r.Route("/course", func(r2 chi.Router) {

			r2.Get("/list", controllerAdmin.AdminGetCourses)
			r2.Get("/list/{id}", controllerAdmin.AdminGetCourseById)
			r2.Post("/create", controllerAdmin.AdminCreateCourse)
			r2.Put("/update", controllerAdmin.AdminUpdateCourse)
			r2.Delete("/delete", controllerAdmin.AdminDeleteCourse)
		})
	})

	app.Route("/teacher", func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middleware.AuthenticationTeacher)

		r.Route("/courses", func(r2 chi.Router) {
			r2.Get("/", controllerTeacher.TeacherGetListCourse)
			r2.Get("/{id}", controllerTeacher.TeacherGetCourseById)

			r2.Route("/{id}/file", func(r3 chi.Router) {

				r3.Use(middleware.CheckCourseOfTeacher)

				r3.Get("/", controllerTeacher.TeacherGetFileOfCourse)
				r3.Get("/{file_id}", controllerTeacher.TeacherOneGetFileOfCourse)
				r3.Post("/create", controllerTeacher.TeacherCreateFileForCourse)
				r3.Put("/update", controllerTeacher.TeacherUpdateFile)
				r3.Delete("/delete", controllerTeacher.TeacherDeleteFileOfCourse)
			})

			r2.Route("/{id}/exercise", func(r3 chi.Router) {

				r3.Use(middleware.CheckCourseOfTeacher)

				r2.Get("/", controllerTeacher.TeacherGetAllExerciseOfCourse)
				r2.Get("/{exercise_id}", controllerTeacher.TeacherGetExerciseById)
				r2.Post("/create", controllerTeacher.TeacherCreateExercise)
				r2.Put("/update", controllerTeacher.TeacherUpdateExercise)
				r2.Delete("/delete", controllerTeacher.TeacherDeleteExercise)
			})

			r2.Route("/{id}/chapter", func(r3 chi.Router) {

				r3.Use(middleware.CheckCourseOfTeacher)

				r2.Get("/", controllerTeacher.TeacherGetAllChapterInCourse)
				r2.Get("/{chapter_id}", controllerTeacher.TeacherGetChapterIdInCourse)
				r2.Post("/create", controllerTeacher.TeacherCreateChapter)
				r2.Put("/update", controllerTeacher.TeacherUpdateChapter)
				r2.Delete("/delete", controllerTeacher.TeacherDeleteChapter)
			})

			r2.Route("/{id}/attendance", func(r3 chi.Router) {

				r3.Use(middleware.CheckCourseOfTeacher)

				r3.Put("/{attendance_id}", controllerTeacher.TeacherUpdateAttendanceStatus)
			})
		})
	})

	app.Route("/student", func(r chi.Router) {

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middleware.AuthenticationStudent)

		r.Post("/joincourse", controllerStudent.StudentJoinCourse)

		r.Route("/my_course", func(r2 chi.Router) {
			r2.Get("/", controllerStudent.StudentGetAllCourse)
			r2.Get("/{id}", controllerStudent.StudentGetCourseById)

			r2.Route("/{id}/exercise", func(r3 chi.Router) {
				r3.Get("/", controllerStudent.StudentAllGetExerciseofCourse)
				r3.Get("/{exercise_id}", controllerStudent.StudentGetOneExerciseOfCourse)
			})
		})

	})

	app.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	port := fmt.Sprintf(":%s", (os.Getenv("PORT")))

	http.ListenAndServe(port, app)
}
