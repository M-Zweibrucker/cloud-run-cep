package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCEPService struct {
	mock.Mock
}

func (m *MockCEPService) GetLocationByCEP(cep string) (string, error) {
	args := m.Called(cep)
	return args.String(0), args.Error(1)
}

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) GetWeatherByLocation(location string) (float64, float64, float64, error) {
	args := m.Called(location)
	return args.Get(0).(float64), args.Get(1).(float64), args.Get(2).(float64), args.Error(3)
}

func TestGetWeather_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCEPService := new(MockCEPService)
	mockWeatherService := new(MockWeatherService)

	mockCEPService.On("GetLocationByCEP", "12345678").Return("São Paulo", nil)
	mockWeatherService.On("GetWeatherByLocation", "São Paulo").Return(25.0, 77.0, 298.0, nil)

	handler := NewWeatherHandler(mockCEPService, mockWeatherService)

	router := gin.New()
	router.GET("/weather/:zipcode", handler.GetWeather)

	req, _ := http.NewRequest("GET", "/weather/12345678", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"temp_C":25`)
	assert.Contains(t, w.Body.String(), `"temp_F":77`)
	assert.Contains(t, w.Body.String(), `"temp_K":298`)

	mockCEPService.AssertExpectations(t)
	mockWeatherService.AssertExpectations(t)
}

func TestGetWeather_InvalidZipcode(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCEPService := new(MockCEPService)
	mockWeatherService := new(MockWeatherService)

	mockCEPService.On("GetLocationByCEP", "1234567").Return("", errors.New("invalid zipcode"))

	handler := NewWeatherHandler(mockCEPService, mockWeatherService)

	router := gin.New()
	router.GET("/weather/:zipcode", handler.GetWeather)

	req, _ := http.NewRequest("GET", "/weather/1234567", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "invalid zipcode")

	mockCEPService.AssertExpectations(t)
}

func TestGetWeather_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCEPService := new(MockCEPService)
	mockWeatherService := new(MockWeatherService)

	mockCEPService.On("GetLocationByCEP", "99999999").Return("", errors.New("can not find zipcode"))

	handler := NewWeatherHandler(mockCEPService, mockWeatherService)

	router := gin.New()
	router.GET("/weather/:zipcode", handler.GetWeather)

	req, _ := http.NewRequest("GET", "/weather/99999999", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "can not find zipcode")

	mockCEPService.AssertExpectations(t)
}
