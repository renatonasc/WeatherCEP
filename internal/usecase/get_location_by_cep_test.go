package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCEPService struct{}

func (m *MockCEPService) GetLocationByCep(cep string) (string, error) {
	if cep == "51020220" {
		return "Recife", nil
	}
	return "", errors.New("CEP not found")
}

func TestGetLocationByCEP(t *testing.T) {
	cepService := &MockCEPService{}
	useCase := GetLocationByCepUseCase{CepService: cepService}

	cep := "51020220"
	location, err := useCase.Execute(cep)
	assert.NoError(t, err)
	assert.Equal(t, "Recife", location)

	cep = "51020221"
	location, err = useCase.Execute(cep)
	assert.Error(t, err)
	assert.Equal(t, "", location)
}
