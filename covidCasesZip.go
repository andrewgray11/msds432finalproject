package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)
const (
	url    = "https://data.cityofchicago.org/resource/yhhz-zm2v.json?$limit=500"
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "postgres"
	dbName = "testdb"
)
type CovidZip struct {
	Zip          string  `json:"zip_code"`
	WeekNumber   int     `json:"week_number"`
	WeekStart    string  `json:"week_start"`
	WeekEnd      string  `json:"week_end"`
	CasesWeekly  int     `json:"cases_weekly"`
	CasesMonthly int     `json:"cases_monthly"`
	CasesCum     int     `json:"cases_cumulative"`
	CaseRateWkly float64 `json:"case_rate_weekly"`
	CaseRateMnth float64 `json:"case_rate_monthly"`
	CaseRateCum  float64 `json:"case_rate_cumulative"`
	TestsWeekly  int     `json:"tests_weekly"`
	TestsCum     int     `json:"tests_cumulative"`
	TestRateWkly float64 `json:"test_rate_weekly"`
	TestRateCum  float64 `json:"test_rate_cumulative"`
	PosRateWkly  float64 `json:"percent_tested_positive_weekly"`
	PosRateCum   float64 `json:"percent_tested_positive_cumulative"`
	DeathsWkly   int     `json:"deaths_weekly"`
	DeathsCum    int     `json:"deaths_cumulative"`
	DeathRtWkly  float64 `json:"death_rate_weekly"`
	DeathRtCum   float64 `json:"death_rate_cumulative"`
	Population   int     `json:"population"`
	RowId        string  `json:"row_id"`
}

func main() {
	// Connect to PostgreSQL database
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Make the API call and retrieve the response body
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON response into slice of CovidZip structs
	var rows []CovidZip
	err = json.Unmarshal(body, &rows)
	if err != nil {
		panic(err)
	}

	// Insert data into database
	ffor _, row := range rows {
		_, err := db.Exec(`
			INSERT INTO covid_zip (
				zip_code, 
				week_number, 
				week_start, 
				week_end, 
				cases_weekly, 
				cases_monthly, 
				cases_cumulative, 
				case_rate_weekly, 
				case_rate_monthly, 
				case_rate_cumulative, 
				tests_weekly, 
				tests_cumulative, 
				test_rate_weekly, 
				test_rate_cumulative, 
				percent_tested_positive_weekly, 
				percent_tested_positive_cumulative, 
				deaths_weekly, 
				deaths_cumulative, 
				death_rate_weekly, 
				death_rate_cumulative, 
				population, 
				row_id
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		`, row.Zip, row.WeekNumber, row.WeekStart, row.WeekEnd, row.CasesWeekly, row.CasesMonthly, row.CasesCum, row.CaseRateWkly, row.CaseRateMnth, row.CaseRateCum, row.TestsWeekly, row.TestsCum, row.TestRateWkly, row.TestRateCum, row.PosRateWkly, row.PosRateCum, row.DeathsWkly, row.DeathsCum, row.DeathRtWkly, row.DeathRtCum, row.Population, row.RowId)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Data has been successfully inserted into the database.")
}
