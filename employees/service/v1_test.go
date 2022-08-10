package service_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestIndex(t *testing.T) {
	// repo := repository.NewInMem()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	mockRepo.EXPECT().ListAll().Return([]entities.Employee{}, nil)

	sut := service.NewV1(mockRepo)

	emps, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, emps)
}
