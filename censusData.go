package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	url    = "https://data.cityofchicago.org/resource/kn9c-c2s2.json"
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "postgres"
	dbName = "testdb"
)

type SocioeconomicData struct {
	CommunityAreaNumber                   int     `json:"community_area_number"`
	CommunityAreaName                     string  `json:"community_area_name"`
	PercentOfHousingCrowded		      float64 `json:"percent_of_housing_crowded"`
	PercentHouseholdsBelowPoverty         float64 `json:"percent_households_below_poverty"`
	PercentAged16Unemployed               float64 `json:"percent_aged_16_unemployed"`
	PercentAged25WithoutHighSchoolDiploma float64 `json:"percent_aged_25_without_high_school_diploma"`
	PercentAgedUnder18OrOver64            float64 `json:"percent_aged_under_18_or_over_64"`
	PerCapitaIncome                       float64 `json:"per_capita_income"`
	HardshipIndex                         int     `json:"hardship_index"`
}

func main() {
	// Make the API call and retrieve the response body
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the response body into an array of SocioeconomicData structs
	var data []SocioeconomicData
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Open a connection to the PostgreSQL database
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert each row into the database
	for _, d := range data {
		query := `INSERT INTO socioeconomic_data (community_area_number, community_area_name, percent_of_housing_crowded, percent_households_below_poverty,
			percent_aged_16_unemployed, percent_aged_25_without_high_school_diploma, percent_aged_under_18_or_over_64,
			per_capita_income, hardship_index)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err := db.Exec(query, d.CommunityAreaNumber, d.CommunityAreaName, d.PercentOfHousingCrowded, d.PercentHouseholdsBelowPoverty,
			d.PercentAged16Unemployed, d.PercentAged25WithoutHighSchoolDiploma, d.PercentAgedUnder18OrOver64,
			d.PerCapitaIncome, d.HardshipIndex)
		if err != nil {
			panic(err)
		}
	}
}
