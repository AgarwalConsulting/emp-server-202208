package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	emps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, emps)

	initialCount := len(emps)

	assert.Equal(t, 3, initialCount)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			emp := entities.Employee{Name: "Gaurav", Department: "LnD"}
			_, err := sut.Save(emp)
			assert.Nil(t, err)
			_, err = sut.ListAll()
			assert.Nil(t, err)
			_, err = sut.FindBy(1)
			assert.Nil(t, err)
		}()
	}

	wg.Wait()

	emps, err = sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, emps)

	finalCount := len(emps)

	assert.Equal(t, 103, finalCount)
}
