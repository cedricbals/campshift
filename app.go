package main

import (
	"github.com/kataras/iris"
	"github.com/joho/godotenv"
	"log"
	"github.com/briandowns/openweathermap"
	"time"
)

type viewDataHolder struct {
	Data         []Day
	Weather      *openweathermap.CurrentWeatherData
	WeatherImage string
}

var viewData viewDataHolder

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file not found")
	}
	
	app := iris.New()
	
	app.StaticEmbeddedGzip("/assets", "./assets", GzipAsset, GzipAssetNames)
	tmpl := iris.HTML("./templates", ".html").Reload(true)
	tmpl.Layout("layout.html")
	tmpl.Binary(Asset, AssetNames)
	app.RegisterView(tmpl)
	
	getViewData()
	
	go refreshViewData()
	
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("", viewData)
		ctx.View("index.html")
	})
	
	app.Run(iris.Addr(":8080"))
}

func getViewData() {
	println("fetched view data")
	
	data := fetchSheetsData()
	
	viewData = viewDataHolder{
		Data:         data,
	}
}

func refreshViewData() {
	for range time.Tick(5 * time.Minute) {
		getViewData()
	}
}
