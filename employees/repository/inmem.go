package repository

import (
	"errors"
	"sync"

	"algogrit.com/emp-server/entities"
)

type inmem struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inmem) ListAll() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()
	return repo.employees, nil
}

func (repo *inmem) Save(newEmployee entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()
	newEmployee.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmployee)

	return &newEmployee, nil
}

func (repo *inmem) FindBy(empID int) (*entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

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
