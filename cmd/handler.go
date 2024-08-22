package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/appnaconda/weather"
)

func ForecastHandler(provider weather.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		if err != nil {
			http.Error(w, "invalid latitude", http.StatusBadRequest)
			return
		}

		lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
		if err != nil {
			http.Error(w, "invalid longitude", http.StatusBadRequest)
			return
		}

		f, err := provider.Forecast(lat, lon)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get forecast: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal(f)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}
