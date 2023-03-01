package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	url    = "https://data.cityofchicago.org/resource/xhc6-88s9.json"
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "postgres"
	dbName = "testdb"
)

type CCVRow struct {
	Index                int     `json:"index_"`
	CommunityArea        string  `json:"community_area"`
	OverallScore         float64 `json:"overall_score"`
	SocioeconomicScore   float64 `json:"socioeconomic_score"`
	HouseholdCrowding    float64 `json:"household_crowding"`
	NoVehicleHouseholds  float64 `json:"no_vehicle_households"`
	PerCapitaIncome      float64 `json:"per_capita_income"`
	Unemployment         float64 `json:"unemployment"`
	NoHighSchoolDiploma  float64 `json:"no_high_school_diploma"`
	AgeAdjustedDeathRate float64 `json:"age_adjusted_death_rate"`
	DiabetesPrevalence   float64 `json:"diabetes_prevalence"`
	HIVPrevalenceRate    float64 `json:"hiv_prevalence_rate"`
	InfantMortalityRate  float64 `json:"infant_mortality_rate"`
	LeadPoisoningRate    float64 `json:"lead_poisoning_rate"`
	HospitalizationRate  float64 `json:"hospitalization_rate"`
}

func main() {
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

	// Unmarshal the response body into an array of CCVRow structs
	var rows []CCVRow
	err = json.Unmarshal(body, &rows)
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
	for _, row := range rows {
		query := `INSERT INTO ccv_data (index_, community_area, overall_score, socioeconomic_score, household_crowding,
                   no_vehicle_households, per_capita_income, unemployment, no_high_school_diploma,
                   age_adjusted_death_rate, diabetes_prevalence, hiv_prevalence_rate, infant_mortality_rate,
                   lead_poisoning_rate, hospitalization_rate)
                  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
		_, err := db.Exec(query, row.Index, row.CommunityArea, row.OverallScore, row.SocioeconomicScore,
			row.HouseholdCrowding, row.NoVehicleHouseholds, row.PerCapitaIncome, row.Unemployment,
			row.NoHighSchoolDiploma, row.AgeAdjustedDeathRate, row.DiabetesPrevalence, row.HIVPrevalenceRate, row.InfantMortalityRate, row.LeadPoisoningRate, row.HospitalizationRate)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Data has been successfully inserted into the database.")
}