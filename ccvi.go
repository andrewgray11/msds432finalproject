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
	GeographyType        string  `json:"geography_type"`
	CommunityAreaOrZip   int     `json:"community_area_or_zip"`
	CommunityAreaName    string  `json:"community_area_name"`
	CcviCategory         string  `json:"ccvi_category"`
	CcviScore            float64 `json:"ccvi_score"`
	SocioeconomicStatus  int     `json:"rank_socioeconomic_status"`
	HouseholdComposition int     `json:"rank_household_composition"`
	NoPrimaryCare        int     `json:"rank_adults_no_pcp"`
	CumMobilityRatio     int     `json:"rank_cumulative_mobility_ratio"`
	FrontlineWorkers     int     `json:"rank_frontline_essential_workers"`
	Age65OrGreater       int     `json:"rank_age_65_plus"`
	ComorbidConditions   int     `json:"rank_comorbid_conditions"`
	CovidIncidenceRate   int     `json:"rank_covid_19_incidence_rate"`
	CovidHospitalRate    int     `json:"rank_covid_19_hospital_admission_rate"`
	CrudeMortalityRate   int     `json:"rank_covid_19_hospital_admission_rate"`
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
	for _, row := range ccvRows {
		query := `INSERT INTO ccv_data (geography_type, community_area_or_zip, community_area_name, ccvi_category, ccvi_score,
					rank_socioeconomic_status, rank_household_composition, rank_adults_no_pcp, rank_cumulative_mobility_ratio,
					rank_frontline_essential_workers, rank_age_65_plus, rank_comorbid_conditions, rank_covid_19_incidence_rate,
					rank_covid_19_hospital_admission_rate, rank_covid_19_hospital_admission_rate)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
		_, err := db.Exec(query, row.GeographyType, row.CommunityAreaOrZip, row.CommunityAreaName, row.CcviCategory,
			row.CcviScore, row.SocioeconomicStatus, row.HouseholdComposition, row.NoPrimaryCare,
			row.CumMobilityRatio, row.FrontlineWorkers, row.Age65OrGreater, row.ComorbidConditions,
			row.CovidIncidenceRate, row.CovidHospitalRate, row.CrudeMortalityRate)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Data has been successfully inserted into the database.")

}
