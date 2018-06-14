package main

import (
	"os"
	"log"
	"strconv"
	owm "github.com/briandowns/openweathermap"
)

func getWeather() (*owm.CurrentWeatherData, string) {
	w, err := owm.NewCurrent("C", "DE", os.Getenv("OWM_API_KEY"))
	if err != nil {
		log.Fatalln(err)
	}
	latitude, err := strconv.ParseFloat(os.Getenv("OWM_LATITUDE"), 64)
	if err != nil {
		log.Fatalln(err)
	}
	longitude, err := strconv.ParseFloat(os.Getenv("OWM_LONGITUDE"), 64)
	if err != nil {
		log.Fatalln(err)
	}
	
	w.CurrentByCoordinates(&owm.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	})
	
	weatherImage := "sun.svg"
	switch w.Weather[0].Main {
	case "Clouds":
		if w.Weather[0].ID == 801 {
			weatherImage = "partly-cloudy.svg"
		} else {
			weatherImage = "cloudy.svg"
		}
	case "Rain":
		weatherImage = "rain.svg"
	case "Drizzle":
		weatherImage = "rain.svg"
	case "Thunderstorm":
		weatherImage = "thunder.svg"
	case "Snow":
		weatherImage = "rain.svg"
	}
	
	return w, weatherImage
}