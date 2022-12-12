package login

import (
	errRequest "app/errrequest"
	"app/model"
	checkexist "app/repository/checkexist"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

// @Summary
// @Description Admin login
// @Tags Login
// @Accept json
// @Produce json
// @Param req body model.Admin false "Account admin" Format(model.Admin)
// @Router /login/admin [post]
func AdminLogin(w http.ResponseWriter, r *http.Request) {

	var admin model.Admin

	json.NewDecoder(r.Body).Decode(&admin)

	accountAdmin, err := checkexist.CheckexistAccountAdmin(admin)

	if err != nil {
		errRequest.NotFoundRequest(w, r, err, "Not found account")
	} else {

		tokenAuth := jwtauth.New("HS256", []byte("token"), nil)
		_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
			"id":         accountAdmin.Id,
			"name":       accountAdmin.Name,
			"permission": "admin",
		})

		if err != nil {
			errRequest.ErrToken(w, r, err)
		} else {

			cookie := http.Cookie{
				Name:     "jwt",
				Value:    tokenString,
				Path:     "/admin",
				HttpOnly: true,
				MaxAge:   60 * 60 * 3,
			}

			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					Token: tokenString,
				},
				Message: "",
				Err:     nil,
				Success: true,
			})

			http.SetCookie(w, &cookie)
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}

	}
}

// @Summary
// @Description
// @Tags Login
// @Accept json
// @Produce json
// @Param req body model.Teacher false "Account Teaccher" Format(model.Teacher)
// @Router /login/teacher [post]
func LoginTeacher(w http.ResponseWriter, r *http.Request) {

	var teacher model.Teacher
	json.NewDecoder(r.Body).Decode(&teacher)

	accountTeacher, err := checkexist.CheckexistAccounTeacher(teacher)

	if err != nil {
		errRequest.NotFoundRequest(w, r, err, "Not exist account")
	} else {

		tokenAuth := jwtauth.New("HS256", []byte("token"), nil)

		_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
			"id":         accountTeacher.Id,
			"name":       accountTeacher.Name,
			"permission": "teacher",
		})

		if err != nil {
			errRequest.ErrToken(w, r, err)
		} else {

			cookie := http.Cookie{
				Name:     "jwt",
				Value:    tokenString,
				Path:     "/teacher",
				HttpOnly: true,
				MaxAge:   60 * 60 * 3,
			}

			http.SetCookie(w, &cookie)

			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					Token: tokenString,
				},
				Message: "Login Successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Login student
// @Tags Login
// @Accept json
// @Produce json
// @Param req body model.Student false "Account student" Format(model.Student)
// @Router /login/student [post]
func LoginStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)

	accountStuddent, err := checkexist.CheckexistAccountStudent(student)

	if err != nil {
		errRequest.NotFoundRequest(w, r, err, "Not found account")
	} else {

		tokenAuth := jwtauth.New("HS256", []byte("token"), nil)
		_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
			"id":         accountStuddent.Id,
			"name":       accountStuddent.Name,
			"permission": "student",
		})

		if err != nil {
			errRequest.ErrToken(w, r, err)
		} else {

			cookie := http.Cookie{
				Name:     "jwt",
				Value:    tokenString,
				Path:     "/student",
				HttpOnly: true,
				MaxAge:   60 * 60 * 3,
			}

			http.SetCookie(w, &cookie)

			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					Token: tokenString,
				},
				Message: "Login Successfully",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}
