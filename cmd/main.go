package main

import (
	"log"
	"net/http"
	"time"

	"github.com/appnaconda/weather/provider/nws"
)

func main() {
	weatherProvider := nws.New(&http.Client{
		Timeout: 5 * time.Second,
	})

	http.HandleFunc("/forecast", ForecastHandler(weatherProvider))

	log.Println("starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
