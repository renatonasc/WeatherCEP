package usecase

import (
	"log"
	"renatonasc/weathercep/internal/entity"
)

type GetWeatherByCepUseCase struct {
}

func (u *GetWeatherByCepUseCase) Execute(cep string) (*entity.WeaterRespose, error) {

	getLocationUseCase := GetLocationByCepUseCase{}
	location, err := getLocationUseCase.Execute(cep)
	log.Println("Location: ", location)
	log.Println("Error: ", err)
	if err != nil {
		return nil, err
	}

	getWeatherByLocationUseCase := GetWeatherByLocationUsecase{}
	weather, err := getWeatherByLocationUseCase.Execute(location)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
