package postgres

import (
	"errors"
	"regexp"
	"strings"
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

func (e *Employee) Update() error {
	db := GetConnection()
	model :=new(Employee)
	err := db.Model(model).Where("id = ?", e.Id).Select()
	if err != nil {
		return err
	}
	e.Created = model.Created
	e.Updated = time.Now()
	return db.Update(e)
}

func (e *Employee) Delete() error {
	db := GetConnection()
	return db.Delete(e)
}

func SelectEmployeeRest() (*[]EmployeeRest, error) {
	db := GetConnection()
	employee := new([]EmployeeRest)
	_, err := db.Query(employee, "SELECT e.id, e.name, e.email, d.name as department from employee e INNER JOIN department d ON e.department=d.id order by e.id")
	return employee, err
}

var reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (e *Employee) Validate() error {
	if e.Department == 0 {
		return errors.New("select one department")
	}
	if e.Name == "" || strings.TrimSpace(e.Name) == "" {
		return errors.New("name is empty")
	}
	if !reEmail.MatchString(e.Email) {
		return errors.New("e-mail invalid")
	}
	return nil
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

func (e *Department) Update() error {
	db := GetConnection()
	return db.Update(e)
}

func (e *Department) Delete() error {
	db := GetConnection()
	return db.Delete(e)
}

func (d *Department) Validate() error {
	if d.Name == "" || strings.TrimSpace(d.Name) == "" {
		return errors.New("name is empty")
	}
	return nil
}