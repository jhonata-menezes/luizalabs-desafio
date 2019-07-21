package postgres

import (
	"time"
)

type Employee struct {
	TableName 	struct{} `sql:"employee" json:"-"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Department int64 `json:"department"`
	Created time.Time `json:"-"`
	Updated time.Time `json:"-"`
}

type EmployeeRest struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Department string `json:"department"`
}

func (e *Employee) Insert() error {
	db := GetConnection()
	e.Created = time.Now()
	e.Updated = time.Now()
	return db.Insert(e)
}
func (e *Employee) Delete() error {
	db := GetConnection()
	return db.Delete(e)
}

func SelectEmployeeRest() (*[]EmployeeRest, error) {
	db := GetConnection()
	employee := new([]EmployeeRest)
	_, err := db.Query(employee, "SELECT e.id, e.name, e.email, d.name as department from employee e INNER JOIN department d ON e.department=d.id")
	return employee, err
}

type Department struct {
	TableName 	struct{} `sql:"department" json:"-"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}

func SelectDepartmentRest() (*[]Department, error) {
	db := GetConnection()
	department := new([]Department)
	err := db.Model(department).Select()
	return department, err
}

func (e *Department) Insert() error {
	db := GetConnection()
	return db.Insert(e)
}

func (e *Department) Delete() error {
	db := GetConnection()
	return db.Delete(e)
}