package admin

import (
	errRequest "app/errrequest"
	"app/model"
	"app/repository/checkexist"
	getData "app/repository/getdata"
	moduleAdmin "app/repository/module/admin"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Admin create account teacher
// @Tags Admin with account teacher
// @Accept json
// @Produce json
// @Param req body model.Teacher false "Account teacher" Format(model.Teacher)
// @Router /admin/teacher/create [post]
func AdminCreateAccountAdmin(w http.ResponseWriter, r *http.Request) {

	var teacher model.Teacher

	json.NewDecoder(r.Body).Decode(&teacher)

	statusEmail := checkexist.CheckexistEmailTeacher(teacher)

	if statusEmail == nil {
		errRequest.ExistEmailRequest(w, r, "Exist email teacher")
	} else {

		err := moduleAdmin.AdminCreateAccountTeacher(teacher)

		if err != nil {
			errRequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Created Successfully!",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Admin get all account teacher
// @Tags Admin with account teacher
// @Accept json
// @Produce json
// @Router /admin/teacher/accounts [get]
func AdminGetAccountTeachers(w http.ResponseWriter, r *http.Request) {

	listAccountTeacher, err := getData.GetAllAccountTeacher()

	if err != nil {
		errRequest.BadRequest(w, r, err)
	} else {

		result, _ := json.Marshal(model.Request{
			Data: model.DataRequest{
				ListAccountTeacher: listAccountTeacher,
			},
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}

}

// @Summary
// @Description Admin get account teacher by id
// @Tags  Admin with account teacher
// @Accept json
// @Produce json
// @Param id path string true "Teacher Id"
// @Router /admin/teacher/accounts/{id} [get]
func AdminGetAccountTeacherById(w http.ResponseWriter, r *http.Request) {

	paramTeacherId := chi.URLParam(r, "id")
	teacherId, err := strconv.Atoi(paramTeacherId)

	if err != nil {
		errRequest.BadRequest(w, r, err)
	} else {

		accountTeacher, err := getData.GetAccountTeacherById(teacherId)

		if err != nil {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Not exist account teacher",
				Err:     err,
				Success: false,
			})
			errRequest.NotFoundRequest(w, r, err, "Not exist account teacher")
			w.Write(result)
		} else {

			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					AccountTeacherId: accountTeacher,
				},
				Message: "",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Desrciption Admin update account teacher
// @Tags Admin with account teacher
// @Accept json
// @Produce json
// @Param req body model.Teacher false "Account teacher need update" Format(model.Teacher)
// @Router /admin/teacher/update [put]
func AdminUpdateAccountTeacher(w http.ResponseWriter, r *http.Request) {

	var accountTeacher model.Teacher

	json.NewDecoder(r.Body).Decode(&accountTeacher)

	err := moduleAdmin.AdminUpdateAccountTeacher(accountTeacher)

	if err != nil {
		errRequest.NotSuccess(w, r, err)
	} else {

		result, _ := json.Marshal(model.Request{
			Data:    nil,
			Message: "Update Successfully",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

// @Summary
// @Description Admin delete account teacher
// @Tags Admin with account teacher
// @Accept json
// @Produde json
// @Param req body model.Teacher false "Account teacher need delete" Format(model.Teacher)
// @Router /admin/teacher/delete [delete]
func AdminDeleteAccountTeacher(w http.ResponseWriter, r *http.Request) {

	var accountTeacher model.Teacher

	json.NewDecoder(r.Body).Decode(&accountTeacher)

	err := moduleAdmin.AdminDeleteAccountTeacher(accountTeacher)

	if err != nil {
		errRequest.NotSuccess(w, r, err)
	} else {

		result, _ := json.Marshal(model.Request{
			Data:    nil,
			Message: "Deleted Successfully!",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
