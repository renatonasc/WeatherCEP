package webclient

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"renatonasc/weathercep/internal/entity"
)

type WeatherService interface {
	GetWeatherByLoctaion(location string) (*entity.WeaterRespose, error)
}

type WeatherAPIDTO struct {
	Current struct {
		Temperature float32 `json:"temp_c"`
	} `json:"current"`
}

type WeatherAPIClient struct {
}

func NewWeatherAPIClient() *WeatherAPIClient {
	return &WeatherAPIClient{}
}

func (w *WeatherAPIClient) GetWeatherByLoctaion(location string) (*entity.WeaterRespose, error) {

	context := context.Background()

	encodedString := url.QueryEscape(location)
	url := "https://api.weatherapi.com/v1/current.json?key=a9037ed280dd4f29940155900241807" + "&q=" + encodedString + "&aqi=no"
	req, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer a requisição" + err.Error())
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro ao ler o body" + err.Error())
		return nil, err
	}

	log.Println("Body: ", string(body))
	var weatherAPIDTO WeatherAPIDTO
	err = json.Unmarshal(body, &weatherAPIDTO)
	if err != nil {
		return nil, err
	}

	weather := &entity.WeaterRespose{
		Location:              location,
		Temparatue_celcius:    weatherAPIDTO.Current.Temperature,
		Temparatue_fahrenheit: weatherAPIDTO.Current.Temperature*1.8 + 32,
		Temperature_kelvin:    weatherAPIDTO.Current.Temperature + 273,
	}

	return weather, nil

}
