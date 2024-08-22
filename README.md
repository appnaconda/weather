# Weather Service
This is a simple weather service that provides weather information for a given latitude and longitude. The service uses the National Weather Service API to get the weather information. 


The service provides a single endpoint `/forecast` that accepts a `GET` request with query parameters `lat` and `lon` for latitude and longitude respectively. The service returns the weather information for the given latitude and longitude in JSON format. 

#### Example Request
```bash
curl http://localhost:8080/forecast?lat=27.763590&lon=-82.400307
```

#### Example Response
```json
{
  "TemperatureUnit": "F",
  "Temperature": 88,
  "TemperatureCondition": "hot",
  "ShortForecast": "Scattered Showers And Thunderstorms",
  "City": "Apollo Beach",
  "State": "FL"
}
```

### Running the service
To run the service, you need to have Go installed on your machine. You can download Go from [here](https://golang.org/dl/).

Once you have Go installed, you can run the service using the following commands:
```bash
cd cmd/
go build -o weather && ./weather && rm weather
```

