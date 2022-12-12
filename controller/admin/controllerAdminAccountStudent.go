package admin

import (
	"app/errrequest"
	"app/model"
	"app/repository/checkexist"
	"app/repository/getdata"
	moduleAdmin "app/repository/module/admin"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Summary
// @Description Admin create account student
// @Tags Admin with account student
// @Accept json
// @Produce json
// @Param req body model.Student false "Account student" Format(model.Student)
// @Router /admin/student/create [post]
func AdminCreateAccountStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)

	statusStudent := checkexist.CheckexistEmailStudent(student)

	if statusStudent == nil {
		errrequest.ArealyExist(w, r, "Email student exist")
	} else {
		err := moduleAdmin.AdminCreateAccountStudent(student)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Creates Successfully",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

// @Summary
// @Description Admin get all account student
// @Tags Admin with account student
// @Accept json
// @Produce json
// @Router /admin/student/accounts [get]
func AdminGetAccountStudents(w http.ResponseWriter, r *http.Request) {

	listAccountStudent, err := getdata.GetAllAccountStudent()

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		result, _ := json.Marshal(model.Request{
			Data: model.DataRequest{
				ListAccountStudent: listAccountStudent,
			},
			Message: "",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}

// @Summary
// @Description Admin get account student by id
// @Tags Admin with account student
// @Accept json
// @Produce json
// @Param id path string true "Student id"
// @Router /admin/student/accounts/{id} [get]
func AdminGetAccountStudentById(w http.ResponseWriter, r *http.Request) {

	paramStudentId := chi.URLParam(r, "id")
	studentId, err := strconv.Atoi(paramStudentId)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		accountStudent, err := getdata.GetAccountStudentById(studentId)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {
			result, _ := json.Marshal(model.Request{
				Data: model.DataRequest{
					AccountStudentId: accountStudent,
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
// @Description Admin update account student
// @Tags Admin with account student
// @Accept json
// @Produce json
// @Param req body model.Student false "Account student need update" Format(model.Student)
// @Router /admin/student/update [put]
func AdminUpdateAccountStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)

	statusStudent := checkexist.CheckexistEmailStudent(student)

	if statusStudent == nil {
		errrequest.ArealyExist(w, r, "Email exist")
	} else {
		err := moduleAdmin.AdminUpdateAccountStudent(student)

		if err != nil {
			errrequest.BadRequest(w, r, err)
		} else {

			result, _ := json.Marshal(model.Request{
				Data:    nil,
				Message: "Updated successfully",
				Err:     nil,
				Success: true,
			})

			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}

}

// @Sumamry
// @Description Admin delete account student
// @Tags Admin with account student
// @Accept json
// @Produce json
// @Param req body model.Student false "Account stundent need delete" Format(model.Student)
// @Router /admin/student/delete [delete]
func AdminDeleteAccountStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student
	json.NewDecoder(r.Body).Decode(&student)

	err := moduleAdmin.AdminDeleteAccountStudent(student)

	if err != nil {
		errrequest.BadRequest(w, r, err)
	} else {
		result, _ := json.Marshal(model.Request{
			Data:    nil,
			Message: "Deleted successfully",
			Err:     nil,
			Success: true,
		})

		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
