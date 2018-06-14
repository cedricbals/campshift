package main

import (
	"github.com/kataras/iris"
	"github.com/joho/godotenv"
	"log"
	"github.com/briandowns/openweathermap"
)

type viewData struct {
	Data         []Day
	Weather      *openweathermap.CurrentWeatherData
	WeatherImage string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file not found")
	}
	
	app := iris.New()
	
	app.StaticEmbeddedGzip("/assets", "./assets", GzipAsset, GzipAssetNames)
	tmpl := iris.HTML("./views", ".html").Reload(true)
	tmpl.Layout("layout.html")
	app.RegisterView(tmpl)
	
	viewData := getViewData()
	
	app.Get("/", func(ctx iris.Context) {
		
		ctx.ViewData("", viewData)
		ctx.View("index.html")
	})
	
	app.Run(iris.Addr(":8080"))
}

func getViewData() viewData {
	weatherMaxTemp, weatherImage := getWeather()
	data := fetchSheetsData()
	
	return viewData{
		Data:         data,
		Weather:      weatherMaxTemp,
		WeatherImage: weatherImage,
	}
}
