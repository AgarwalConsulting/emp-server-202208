package repository

import (
	"errors"

	"algogrit.com/emp-server/entities"
)

type inmem struct {
	employees []entities.Employee
}

func (repo *inmem) ListAll() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmem) Save(newEmployee entities.Employee) (*entities.Employee, error) {
	newEmployee.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmployee)

	return &newEmployee, nil
}

func (repo *inmem) FindBy(empID int) (*entities.Employee, error) {
	if empID >= len(repo.employees) {
		return nil, errors.New("unknown employee")
	}

	return &repo.employees[empID-1], nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Jose", "Cloud", 1002},
		{3, "Prabhakar", "SRE", 10003},
	}

	return &inmem{employees: employees}
}
