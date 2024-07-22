package usecase

import (
	"renatonasc/weathercep/internal/entity"
	"renatonasc/weathercep/internal/infra/web/webclient"
)

type GetWeatherByLocationUsecase struct {
	WeatherService webclient.WeatherService
}

func (u *GetWeatherByLocationUsecase) Execute(location string) (*entity.WeaterRespose, error) {

	weather, err := u.WeatherService.GetWeatherByLoctaion(location) //   webclient.NewWeatherAPIClient().GetWeatherByLoctaion(location)
	if err != nil {
		return nil, err
	}

	return weather, nil

}
