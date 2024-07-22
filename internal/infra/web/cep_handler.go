package web

import (
	"encoding/json"
	"log"
	"net/http"
	"renatonasc/weathercep/internal/infra/web/webclient"
	"renatonasc/weathercep/internal/usecase"

	"github.com/go-chi/chi/v5"
)

type Error struct {
	Message string `json:"message"`
}

type CepHandler struct {
}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

func (h *CepHandler) GetWeatherByCep(w http.ResponseWriter, r *http.Request) {

	cep := chi.URLParam(r, "cep")
	if cep == "" {
		http.Error(w, "File cep is required", http.StatusBadRequest)
		return
	}

	log.Println("CEP informado: ", cep)
	getWeatherByCepUsecase := usecase.GetWeatherByCepUseCase{
		CepService:     webclient.NewViaCepClient(),
		WeatherService: webclient.NewWeatherAPIClient(),
	}
	weather, err := getWeatherByCepUsecase.Execute(cep)
	if err != nil {
		if err.Error() == "CEP deve conter 8 digitos" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if err.Error() == "CEP não encontrado" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
