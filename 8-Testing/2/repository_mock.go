package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRespositoryMock struct {
	mock.Mock
}

func (m *TaxRespositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}