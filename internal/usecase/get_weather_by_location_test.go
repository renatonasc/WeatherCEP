package usecase

import (
	"errors"
	"renatonasc/weathercep/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockWeatherService struct{}

func (m *MockWeatherService) GetWeatherByLoctaion(location string) (*entity.WeaterRespose, error) {
	if location == "Recife" {
		return &entity.WeaterRespose{
			Location:              location,
			Temparatue_celcius:    30.0,
			Temparatue_fahrenheit: 86.0,
			Temperature_kelvin:    303.15,
		}, nil
	}
	return nil, errors.New("Location not found")
}

func TestGetWeatherByLocation(t *testing.T) {
	weatherService := &MockWeatherService{}
	useCase := GetWeatherByLocationUsecase{WeatherService: weatherService}

	location := "Recife"
	weather, err := useCase.Execute(location)
	assert.NoError(t, err)
	assert.Equal(t, float32(30.0), weather.Temparatue_celcius)

	location = "Sao Paulo"
	weather, err = useCase.Execute(location)
	assert.Error(t, err)
	assert.Nil(t, weather)
}
