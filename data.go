package main

import (
	"log"
	"strings"
	"os"
	"time"
)

type groupedEntryMap map[string][]Entry

type Day struct {
	PreviousDate string
	NextDate string
	
	Name string
	Date string
	
	Entries       *groupedEntryMap
	GroupingOrder *[]string
	IsToday bool
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
	
	todayInSheetFormat := time.Now().Format("02.01.2006")
	foundToday := false
	
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
			IsToday: todayInSheetFormat == dayNameAndDate[1],
		})
		
		if !foundToday {
			foundToday = todayInSheetFormat == dayNameAndDate[1]
		}
		
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
	
	if !foundToday {
		data[0].IsToday = true
	}
	
	// build "linked list"
	var lastDate string
	for dayIdx, day := range data {
		if len(lastDate) > 0 {
			data[dayIdx].PreviousDate = lastDate
			data[dayIdx - 1].NextDate = day.Date
		}
		lastDate = day.Date
	}
	
	return
}
