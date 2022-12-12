package middleware

import (
	"app/errrequest"
	"app/repository/checkexist"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func AuthenticationTeacher(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, claims, err := jwtauth.FromContext(r.Context())

		if err != nil {
			badRequest(w, r, err)
		} else {
			if claims["permission"] == "teacher" {
				next.ServeHTTP(w, r)
			} else {
				notIsPermission(w, r, "Not is Teacher")
			}
		}
	})
}

func CheckCourseOfTeacher(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		courseId, err1 := strconv.Atoi(chi.URLParam(r, "id"))
		_, claims, err2 := jwtauth.FromContext(r.Context())

		if err1 != nil {
			errrequest.BadRequest(w, r, err1)
		} else if err2 != nil {
			errrequest.ErrToken(w, r, err2)
		} else {
			teacherId := int(claims["id"].(float64))
			statusCourse := checkexist.CheckexistCourseOfTeacher(teacherId, courseId)

			if statusCourse != nil {
				errrequest.BadRequest(w, r, statusCourse)
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}
