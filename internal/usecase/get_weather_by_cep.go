package usecase

import (
	"log"
	"renatonasc/weathercep/internal/entity"
	"renatonasc/weathercep/internal/infra/web/webclient"
)

type GetWeatherByCepUseCase struct {
	CepService     webclient.CepService
	WeatherService webclient.WeatherService
}

func (u *GetWeatherByCepUseCase) Execute(cep string) (*entity.WeaterRespose, error) {

	getLocationUseCase := GetLocationByCepUseCase{u.CepService}
	location, err := getLocationUseCase.Execute(cep)
	log.Println("Location: ", location)
	log.Println("Error: ", err)
	if err != nil {
		return nil, err
	}

	getWeatherByLocationUseCase := GetWeatherByLocationUsecase{u.WeatherService}
	weather, err := getWeatherByLocationUseCase.Execute(location)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
