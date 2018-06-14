package main

import (
	"log"
	"strings"
	"os"
)

type groupedEntryMap map[string][]Entry

type Day struct {
	Name string
	Date string
	
	Entries       *groupedEntryMap
	GroupingOrder *[]string
}

type Entry struct {
	Name               string
	Description        string
	ResponsiblePersons string
	Icon               string
	Color              string
	Row                int
}

func fetchSheetsData() (data []Day) {
	sheetId := os.Getenv("GOOGLE_SHEETS_SPREADSHEET_ID")
	sheetsClient := initGoogleSheetsClient()
	resp, err := sheetsClient.Spreadsheets.Get(sheetId).Do()
	if err != nil {
		log.Fatalln(err)
	}
	days := resp.Sheets
	for _, day := range days {
		dayResp, err := sheetsClient.Spreadsheets.Values.Get(sheetId, "'"+day.Properties.Title+"'!A2:F").Do()
		if err != nil {
			log.Fatalln(err)
		}
		
		entries := groupedEntryMap{}
		var groupingOrder []string
		dayNameAndDate := strings.Split(day.Properties.Title, " ")
		data = append(data, Day{
			Date:          dayNameAndDate[1],
			Name:          dayNameAndDate[0],
			Entries:       &entries,
			GroupingOrder: &groupingOrder,
		})
		
		for rowIdx, row := range dayResp.Values {
			grouping := row[0].(string)
			_, ok := entries[grouping]
			if !ok {
				entries[grouping] = []Entry{}
				groupingOrder = append(groupingOrder, grouping)
			}
			
			entries[grouping] = append(entries[grouping], Entry{
				Name:               row[1].(string),
				Description:        row[2].(string),
				ResponsiblePersons: row[3].(string),
				Color:              row[4].(string),
				Icon:               row[5].(string),
				Row:                rowIdx,
			})
		}
	}
	
	return
}
