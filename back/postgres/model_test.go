package postgres

import "testing"

func TestEmployee(t *testing.T) {

	employee := new(Employee)

	employee.Email = "vemserfeliz@luizalabs.com"
	employee.Name = "Vem Ser Feliz"

	err := employee.Validate()
	if err == nil {
		t.Error("validate ignore department")
	}
	employee.Department = 1
	employee.Email = ""
	err = employee.Validate()
	if err == nil {
		t.Error("validate ignore email")
	}
	employee.Email = "vemserfeliz"
	err = employee.Validate()
	if err == nil {
		t.Error("e-mail not is valid")
	}
}

func TestDepartment(t *testing.T) {
	department := new(Department)
	err := department.Validate()
	if err == nil {
		t.Error("name is empty")
	}
	department.Name = "support"
	err = department.Validate()
	if err != nil {
		t.Error("name is valid", err.Error())
	}
}
