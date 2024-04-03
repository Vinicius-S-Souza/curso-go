package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "amount must be greater then 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateTaxAndSave(t *testing.T) {
	respository := &TaxRespositoryMock{}
	respository.On("SaveTax", 10.0).Return(nil)
	respository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, respository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, respository)
	assert.Error(t, err, "error saving tax")

	respository.AssertExpectations(t)

}
