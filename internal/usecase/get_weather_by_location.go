package usecase

import (
	"renatonasc/weathercep/internal/entity"
	"renatonasc/weathercep/internal/infra/web/webclient"
)

type GetWeatherByLocationUsecase struct {
}

func (u *GetWeatherByLocationUsecase) Execute(location string) (*entity.WeaterRespose, error) {

	weather, err := webclient.NewWeatherAPIClient().GetWeatherByLoctaion(location)
	if err != nil {
		return nil, err
	}

	return weather, nil

}
