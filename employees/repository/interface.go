package repository

import "algogrit.com/emp-server/entities"

//go:generate mockgen -source $GOFILE -destination mock_$GOPACKAGE.go -package $GOPACKAGE

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(entities.Employee) (*entities.Employee, error)
	FindBy(int) (*entities.Employee, error)
}
