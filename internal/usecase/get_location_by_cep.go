package usecase

import (
	"errors"
	"log"
	"regexp"
	"renatonasc/weathercep/internal/infra/web/webclient"
)

type GetLocationByCepUseCase struct {
}

func (u *GetLocationByCepUseCase) Execute(cep string) (string, error) {
	var expReg = regexp.MustCompile(`^\d{5}-?\d{3}$`)

	if !expReg.MatchString(cep) {
		log.Println("CEP deve conter 8 digitos")
		log.Println("CEP informado: ", cep)
		return "", errors.New("CEP deve conter 8 digitos")
	}

	location, err := webclient.NewViaCepClient().GetLocationByCep(cep)
	if err != nil {
		return "", err
	}

	return location, nil
}