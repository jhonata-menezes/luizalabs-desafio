package api

import (
	"github.com/jhonata-menezes/luizalabs-desafio/back/postgres"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func LoadRoutes(e *echo.Echo) {
	e.GET("/employee", GETEmployee)
	e.PUT("/employee", PUTEmployee)
	e.DELETE("/employee/:id", DELETEEmployee)

	//department
	e.GET("/department", GETDepartment)
	e.PUT("/department", PUTDepartment)
	e.DELETE("/department/:id", DELETEDepartment)
}

type ResponseDefault struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func GETEmployee (e echo.Context) error {
	employee, err := postgres.SelectEmployeeRest()
	if err != nil {
		log.Println(err.Error())
		return e.String(http.StatusNoContent, "")
	}
	if len(*employee) == 0 {
		return e.JSON(http.StatusOK, []string{})
	}
	return e.JSON(http.StatusOK, employee)
}

func PUTEmployee (e echo.Context) error {
	employee := new(postgres.Employee)
	err := e.Bind(employee)
	if err != nil {
		return e.JSON(http.StatusOK, &ResponseDefault{"error", err.Error()})
	}
	employee.Email = strings.ToLower(employee.Email)
	if err = employee.Validate(); err != nil {
		return e.JSON(http.StatusOK, &ResponseDefault{"error", err.Error()})
	}
	if employee.Id == 0 {
		err = employee.Insert()
	} else {
		err = employee.Update()
	}
	if err != nil {
		return e.JSON(http.StatusOK, &ResponseDefault{"error", err.Error()})
	}
	return e.JSON(http.StatusOK, &ResponseDefault{
		"ok", "inserted",
	})
}

func DELETEEmployee (e echo.Context) error {
	employee := new(postgres.Employee)
	log.Println("param", e.Param("id"))
	idParam, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusOK, ResponseDefault{"error", "id is not integer"})
	}
	employee.Id = int64(idParam)
	err = employee.Delete()
	if err != nil {
		return e.JSON(http.StatusOK, ResponseDefault{"error", "employee not exist"})
	}
	return e.JSON(http.StatusOK, ResponseDefault{"ok", "deleted"})
}

//department
func GETDepartment (e echo.Context) error {
	department, err := postgres.SelectDepartmentRest()
	if err != nil {
		log.Println(err.Error())
		return e.String(http.StatusNoContent, "")
	}
	if len(*department) == 0 {
		return e.JSON(http.StatusOK, []string{})
	}
	return e.JSON(http.StatusOK, department)
}

func PUTDepartment (e echo.Context) error {
	department := new(postgres.Department)
	err := e.Bind(department)
	if err != nil {
		return e.JSON(http.StatusOK, &ResponseDefault{"error", err.Error()})
	}
	if err = department.Validate(); err != nil {
		return e.JSON(http.StatusOK, &ResponseDefault{"error", err.Error()})
	}
	if department.Id == 0 {
		err = department.Insert()
	} else {
		err = department.Update()
	}
	if err != nil {
		return e.JSON(http.StatusOK, &ResponseDefault{"error", err.Error()})
	}
	return e.JSON(http.StatusOK, &ResponseDefault{
		"ok", "inserted",
	})
}

func DELETEDepartment (e echo.Context) error {
	department := new(postgres.Department)
	idParam, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, ResponseDefault{"error", "id is not integer"})
	}
	department.Id = int64(idParam)
	err = department.Delete()
	if err != nil {
		return e.JSON(http.StatusOK, ResponseDefault{"error", "can not delete"})
	}
	return e.JSON(http.StatusOK, ResponseDefault{"ok", "deleted"})
}

