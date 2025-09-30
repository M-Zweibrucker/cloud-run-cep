package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCEP(t *testing.T) {
	service := &cepService{}

	tests := []struct {
		name     string
		cep      string
		expected error
	}{
		{
			name:     "CEP válido com 8 dígitos",
			cep:      "12345678",
			expected: nil,
		},
		{
			name:     "CEP válido com formatação",
			cep:      "12345-678",
			expected: nil,
		},
		{
			name:     "CEP inválido com menos de 8 dígitos",
			cep:      "1234567",
			expected: assert.AnError,
		},
		{
			name:     "CEP inválido com mais de 8 dígitos",
			cep:      "123456789",
			expected: assert.AnError,
		},
		{
			name:     "CEP inválido com letras",
			cep:      "1234567a",
			expected: assert.AnError,
		},
		{
			name:     "CEP vazio",
			cep:      "",
			expected: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.validateCEP(tt.cep)
			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "0°C para Fahrenheit",
			celsius:  0,
			expected: 32,
		},
		{
			name:     "25°C para Fahrenheit",
			celsius:  25,
			expected: 77,
		},
		{
			name:     "100°C para Fahrenheit",
			celsius:  100,
			expected: 212,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := celsiusToFahrenheit(tt.celsius)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{
			name:     "0°C para Kelvin",
			celsius:  0,
			expected: 273,
		},
		{
			name:     "25°C para Kelvin",
			celsius:  25,
			expected: 298,
		},
		{
			name:     "100°C para Kelvin",
			celsius:  100,
			expected: 373,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := celsiusToKelvin(tt.celsius)
			assert.Equal(t, tt.expected, result)
		})
	}
}
