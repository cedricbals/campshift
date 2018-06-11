package main

import (
	"github.com/kataras/iris"
	"github.com/joho/godotenv"
	"log"
	owm "github.com/briandowns/openweathermap"
	"os"
	"strconv"
	"strings"
)

type viewData struct {
	Data []Day
	Weather float64
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

func getWeather() float64 {
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
	
	return w.Main.TempMax
}

func getViewData() viewData {
	weather := getWeather()
	
	sheetId := os.Getenv("GOOGLE_SHEETS_SPREADSHEET_ID")
	sheetsClient := initGoogleSheetsClient()
	resp, err := sheetsClient.Spreadsheets.Get(sheetId).Do()
	if err != nil {
		log.Fatalln(err)
	}
	days := resp.Sheets
	
	var data []Day
	
	for _, day := range days {
		dayResp, err := sheetsClient.Spreadsheets.Values.Get(sheetId, "'" + day.Properties.Title + "'!A2:E").Do()
		if err != nil {
			log.Fatalln(err)
		}
		
		entries := groupedEntryMap{}
		dayNameAndDate := strings.Split(day.Properties.Title, " ")
		data = append(data, Day{
			Date: dayNameAndDate[1],
			Name: dayNameAndDate[0],
			Entries: &entries,
		})
		
		for _, row := range dayResp.Values {
			grouping := row[0].(string)
			_, ok := entries[grouping]
			if !ok {
				entries[grouping] = []Entry{}
			}
			
			entries[grouping] = append(entries[grouping], Entry{
				Name: row[1].(string),
				Description: row[2].(string),
				ResponsiblePersons: row[3].(string),
				Color: row[4].(string),
			})
		}
	}
	
	return viewData{
		Data: data,
		Weather: weather,
	}
}