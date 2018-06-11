package main

type groupedEntryMap map[string][]Entry

type Day struct {
	Name string
	Date string
	
	Entries *groupedEntryMap
}

type Entry struct {
	Name string
	Description string
	ResponsiblePersons string
	Icon string
	Color string
}