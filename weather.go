package weather

type Forecast struct {
	TemperatureUnit      string
	Temperature          float64
	TemperatureCondition string
	ShortForecast        string
	City                 string
	State                string
}

type Provider interface {
	Forecast(latitude, longitude float64) (*Forecast, error)
}

// TemperatureCondition returns a human-readable description of the temperature
func TemperatureCondition(temp float64) string {
	if temp < 60 {
		return "cold"
	}
	if temp < 80 {
		return "moderate"
	}

	return "hot"
}
