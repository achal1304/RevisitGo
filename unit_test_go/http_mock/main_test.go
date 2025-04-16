package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	fmt.Println(args...)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestFetchWeather(t *testing.T) {
	mockResponse := `{"temperature": 22.5, "humidity": 65}`

	mockClient := &MockClient{}

	mockResp := &http.Response{
		Status:     fmt.Sprint(http.StatusOK),
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader([]byte(mockResponse))),
	}

	mockClient.On("Get", "http://api.weather.com/weather?city=London").Return(mockResp, nil)

	weather, err := FetchWeather("http://api.weather.com", "London", mockClient)
	assert.NoError(t, err)
	assert.NotNil(t, weather)
	assert.Equal(t, 22.5, weather.Temperature)
	assert.Equal(t, 65, weather.Humidity)

	// Verify that the mock was called with the expected URL
	mockClient.AssertExpectations(t)
}
