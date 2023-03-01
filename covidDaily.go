package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type CovidData struct {
	Date                     string `json:"date"`
	TotalCases               string `json:"total_cases"`
	NewCases                 string `json:"new_cases"`
	TotalDeaths              string `json:"total_deaths"`
	NewDeaths                string `json:"new_deaths"`
	TotalHospitalized        string `json:"total_hospitalized"`
	NewHospitalized          string `json:"new_hospitalized"`
	AverageDailyHospitalized string `json:"average_daily_hospitalized"`
}

func main() {
	// URL for the API call
	url := "https://data.cityofchicago.org/resource/naz8-j4nc.json?$limit=500"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Decode the response body into a slice of CovidData structs
	var covidData []CovidData
	err = json.NewDecoder(resp.Body).Decode(&covidData)
	if err != nil {
		panic(err)
	}

	// Connect to the Postgres database
	db, err := sql.Open("postgres", "postgres://user:password@localhost/covid_data")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert each CovidData struct into the database, but only if all columns are not blank
	for _, data := range covidData {
		if data.Date != "" && data.TotalCases != "" && data.NewCases != "" && data.TotalDeaths != "" && data.NewDeaths != "" && data.TotalHospitalized != "" && data.NewHospitalized != "" && data.AverageDailyHospitalized != "" {
			// Prepare the SQL statement
			stmt, err := db.Prepare("INSERT INTO covid_data (date, total_cases, new_cases, total_deaths, new_deaths, total_hospitalized, new_hospitalized, average_daily_hospitalized) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
			if err != nil {
				panic(err)
			}

			// Execute the SQL statement with the values from the CovidData struct
			_, err = stmt.Exec(strings.TrimSpace(data.Date), strings.TrimSpace(data.TotalCases), strings.TrimSpace(data.NewCases), strings.TrimSpace(data.TotalDeaths), strings.TrimSpace(data.NewDeaths), strings.TrimSpace(data.TotalHospitalized), strings.TrimSpace(data.NewHospitalized), strings.TrimSpace(data.AverageDailyHospitalized))
			if err != nil {
				panic(err)
			}
			fmt.Println("Inserted row into database")
		}
	}
}