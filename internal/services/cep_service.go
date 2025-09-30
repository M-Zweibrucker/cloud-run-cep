package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type CEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

type CEPService interface {
	GetLocationByCEP(cep string) (string, error)
}

type cepService struct {
	client *http.Client
}

func NewCEPService() CEPService {
	return &cepService{
		client: &http.Client{},
	}
}

func (s *cepService) validateCEP(cep string) error {
	cleanCEP := regexp.MustCompile(`\D`).ReplaceAllString(cep, "")
	if len(cleanCEP) != 8 {
		return fmt.Errorf("invalid zipcode")
	}

	return nil
}

func (s *cepService) GetLocationByCEP(cep string) (string, error) {
	if err := s.validateCEP(cep); err != nil {
		return "", err
	}

	cleanCEP := regexp.MustCompile(`\D`).ReplaceAllString(cep, "")
	formattedCEP := fmt.Sprintf("%s-%s", cleanCEP[:5], cleanCEP[5:])

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", formattedCEP)
	resp, err := s.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("can not find zipcode")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("can not find zipcode")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("can not find zipcode")
	}

	var cepResp CEPResponse
	if err := json.Unmarshal(body, &cepResp); err != nil {
		return "", fmt.Errorf("can not find zipcode")
	}

	if cepResp.Erro {
		return "", fmt.Errorf("can not find zipcode")
	}

	return cepResp.Localidade, nil
}
