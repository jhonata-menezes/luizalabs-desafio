package api

import (
	"encoding/json"
	"fmt"
	"github.com/jhonata-menezes/luizalabs-desafio/back/postgres"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	employeeBody = postgres.Employee{Name: "Vem Ser Feliz", Department: 1, Email: "vemserfeliz@luizalabs.com"}
	employeeUpdateBody = postgres.Employee{Id: 1, Name: "Vem Ser Felizz", Department: 2, Email: "vemserfelizz@luizalabs.com"}
	employeeCreated = `{"status":"ok","message":"inserted"}` + "\n"
	employeeDeleted = `{"status":"ok","message":"deleted"}` + "\n"
)


func TestCreateUser(t *testing.T) {
	// create department for tests
	InsertDepartment()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/employee", strings.NewReader(StructToJson(employeeBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, PUTEmployee(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, employeeCreated, rec.Body.String())
	}
}

func TestUpdateUser(t *testing.T) {
	employee := new(postgres.Employee)
	db := postgres.GetConnection()
	err := db.Model(employee).Where("email = ?", employeeBody.Email).First()
	assert.Equal(t, err, nil)
	id := employee.Id
	employeeUpdateBody.Id = id

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/employee", strings.NewReader(StructToJson(employeeUpdateBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, PUTEmployee(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, employeeCreated, rec.Body.String())
	}

	err = db.Model(employee).Where("id = ?", id).First()
	assert.Equal(t, err, nil)
	assert.Equal(t, employeeUpdateBody.Id, employee.Id, "employee id not match")
	assert.Equal(t, employeeUpdateBody.Email, employee.Email, "employee not update e-mail")
	assert.Equal(t, employeeUpdateBody.Name, employee.Name, "employee not update name")
	assert.Equal(t, employeeUpdateBody.Department, employee.Department, "employee not update department")
}

func TestDeleteUser(t *testing.T) {
	e := echo.New()
	t.Log(fmt.Sprintf("/employee/%d", employeeUpdateBody.Id))
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(fmt.Sprintf("/employee/%d", employeeUpdateBody.Id))
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", employeeUpdateBody.Id))

	if assert.NoError(t, DELETEEmployee(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, employeeDeleted, rec.Body.String())
	}
}


func StructToJson(v interface{}) string {
	by, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(by)
}

func InsertDepartment() {
	department := new(postgres.Department)
	db := postgres.GetConnection()
	department.Id = 1
	exists, _ := db.Model(department).Exists()
	if !exists {
		department.Name = "support"
		db.Insert(department)
	}

	department.Id = 2
	department.Name = "support 2"
	exists, _ = db.Model(department).Exists()
	if !exists {
		db.Insert(department)
	}

	department.Id = 3
	department.Name = "support 3"
	exists, _ = db.Model(department).Exists()
	if !exists {
		db.Insert(department)
	}
}


