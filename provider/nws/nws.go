package nws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/appnaconda/weather"
)

const baseURL = "https://api.weather.gov"

type (
	// Response is a simplified version of the actual response
	// only containing the properties we need
	Response struct {
		Properties struct {
			ForecastURL string `json:"forecast"`
			Periods     []struct {
				Name            string  `json:"name"`
				Temperature     float64 `json:"temperature"`
				TemperatureUnit string  `json:"temperatureUnit"`
				ShortForecast   string  `json:"shortForecast"`
			} `json:"periods"`
			RelativeLocation struct {
				Properties struct {
					City  string `json:"city"`
					State string `json:"state"`
				}
			} `json:"relativeLocation"`
		}
	}

	ErrorResponse struct {
		CorrelationId string `json:"correlationId"`
		Title         string `json:"title"`
		Type          string `json:"type"`
		Status        int    `json:"status"`
		Detail        string `json:"detail"`
		Instance      string `json:"instance"`
	}
)

// Client is a client for the National Weather Service API that implements the weather.Provider interface
type Client struct {
	http *http.Client
}

func New(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		http: httpClient,
	}
}

func (c *Client) get(url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/geo+json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return nil, fmt.Errorf("error: %s", resp.Status)
		}
		return nil, fmt.Errorf("error: %s", errResp.Detail)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	return &response, err
}

func (c *Client) Forecast(latitude, longitude float64) (*weather.Forecast, error) {
	resp, err := c.get(baseURL + "/points/" + fmt.Sprintf("%f,%f", latitude, longitude))
	if err != nil {
		return nil, fmt.Errorf("error getting forecast: %w", err)
	}

	if resp.Properties.ForecastURL == "" {
		return nil, fmt.Errorf("forecast not found")
	}

	forecast := weather.Forecast{
		City:  resp.Properties.RelativeLocation.Properties.City,
		State: resp.Properties.RelativeLocation.Properties.State,
	}

	resp, err = c.get(resp.Properties.ForecastURL)
	if err != nil {
		return nil, fmt.Errorf("error getting forecast: %w", err)
	}

	if len(resp.Properties.Periods) == 0 {
		return nil, fmt.Errorf("forecast not found")
	}

	period := resp.Properties.Periods[0] // first period is current forecast

	forecast.TemperatureUnit = period.TemperatureUnit
	forecast.Temperature = period.Temperature
	forecast.ShortForecast = period.ShortForecast
	forecast.TemperatureCondition = weather.TemperatureCondition(period.Temperature)

	return &forecast, nil
}
