package repository

import "algogrit.com/emp-server/entities"

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(entities.Employee) (*entities.Employee, error)
	FindBy(int) (*entities.Employee, error)
}
