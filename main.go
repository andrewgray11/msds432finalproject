package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jonas-p/go-shp"
	_ "github.com/lib/pq"
)

type buildingPermit struct {
	ID            string `json:"id"`      // int
	PermitNumber  string `json:"permit_"` // use backtick to handle column name with underscore
	PermitType    string `json:"permit_type"`
	ReviewType    string `json:"review_type"`
	TotalFee      string `json:"total_fee"` //float64
	AppStartDate  string `json:"application_start_date"`
	IssueDate     string `json:"issue_date"`
	CommunityArea string `json:"community_area"` // int
	Latitude      string `json:"latitude"`       //float64
	Longitude     string `json:"longitude"`      //float64
}

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

type SocioeconomicData struct {
	CommunityAreaNumber                   int     `json:"community_area_number"`
	CommunityAreaName                     string  `json:"community_area_name"`
	PercentOfHousingCrowded               float64 `json:"percent_of_housing_crowded"`
	PercentHouseholdsBelowPoverty         float64 `json:"percent_households_below_poverty"`
	PercentAged16Unemployed               float64 `json:"percent_aged_16_unemployed"`
	PercentAged25WithoutHighSchoolDiploma float64 `json:"percent_aged_25_without_high_school_diploma"`
	PercentAgedUnder18OrOver64            float64 `json:"percent_aged_under_18_or_over_64"`
	PerCapitaIncome                       float64 `json:"per_capita_income"`
	HardshipIndex                         int     `json:"hardship_index"`
}

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

type PublicHealthStats struct {
	CommunityArea                         string  `json:"community_area"`
	CommunityAreaName                     string  `json:"community_area_name"`
	BirthRate                             string  `json:"birth_rate"`
	GeneralFertilityRate                  string  `json:"general_fertility_rate"`
	LowBirthWeight                        string  `json:"low_birth_weight"`
	PrenatalCareBeginningInFirstTrimester float64 `json:"prenatal_care_beginning_in_first_trimester"`
	PretermBirths                         string  `json:"preterm_births"`
	TeenBirthRate                         float64 `json:"teen_birth_rate"`
	AssaultHomicide                       string  `json:"assault_homicide"`
	BreastCancerInFemales                 float64 `json:"breast_cancer_in_females"`
	CancerAllSites                        string  `json:"cancer_all_sites"`
	ColorectalCancer                      float64 `json:"colorectal_cancer"`
	DiabetesRelated                       string  `json:"diabetes_related"`
	FirearmRelated                        float64 `json:"firearm_related"`
	InfantMortalityRate                   float64 `json:"infant_mortality_rate"`
	LungCancer                            float64 `json:"lung_cancer"`
	ProstateCancerInMales                 float64 `json:"prostate_cancer_in_males"`
	StrokeCerebrovascularDisease          float64 `json:"stroke_cerebrovascular_disease"`
	ChildhoodBloodLeadLevelScreening      float64 `json:"childhood_blood_lead_level_screening"`
	ChildhoodLeadPoisoning                float64 `json:"childhood_lead_poisoning"`
	GonorrheaInFemales                    float64 `json:"gonorrhea_in_females"`
	GonorrheaInMales                      string  `json:"gonorrhea_in_males"`
	Tuberculosis                          float64 `json:"tuberculosis"`
	BelowPovertyLevel                     float64 `json:"below_poverty_level"`
	CrowdedHousing                        float64 `json:"crowded_housing"`
	Dependency                            float64 `json:"dependency"`
	NoHighSchoolDiploma                   float64 `json:"no_high_school_diploma"`
	PerCapitaIncome                       float64 `json:"per_capita_income"`
	Unemployment                          float64 `json:"unemployment"`
}

type TaxiTrips struct {
	TripID                   string    `json:"trip_id"`
	TaxiID                   string    `json:"taxi_id"`
	TripStartTimestamp       time.Time `json:"trip_start_timestamp"`
	TripEndTimestamp         time.Time `json:"trip_end_timestamp"`
	TripSeconds              int       `json:"trip_seconds"`
	TripMiles                float64   `json:"trip_miles"`
	PickupCensusTract        string    `json:"pickup_census_tract"`
	DropoffCensusTract       string    `json:"dropoff_census_tract"`
	PickupCommunityArea      int       `json:"pickup_community_area"`
	DropoffCommunityArea     int       `json:"dropoff_community_area"`
	Fare                     float64   `json:"fare"`
	Tips                     float64   `json:"tips"`
	Tolls                    float64   `json:"tolls"`
	Extras                   float64   `json:"extras"`
	TripTotal                float64   `json:"trip_total"`
	PaymentType              string    `json:"payment_type"`
	Company                  string    `json:"company"`
	PickupCentroidLatitude   float64   `json:"pickup_centroid_latitude"`
	PickupCentroidLongitude  float64   `json:"pickup_centroid_longitude"`
	PickupCentroidLocation   string    `json:"pickup_centroid_location"`
	DropoffCentroidLatitude  float64   `json:"dropoff_centroid_latitude"`
	DropoffCentroidLongitude float64   `json:"dropoff_centroid_longitude"`
	DropoffCentroidLocation  string    `json:"dropoff_centroid_location"`
}

type TNPTrips struct {
	TripID                   string  `json:"trip_id"`
	TripStartTimestamp       string  `json:"trip_start_timestamp"`
	TripEndTimestamp         string  `json:"trip_end_timestamp"`
	TripSeconds              int     `json:"trip_seconds"`
	TripMiles                float64 `json:"trip_miles"`
	PickupCensusTract        string  `json:"pickup_census_tract"`
	DropoffCensusTract       string  `json:"dropoff_census_tract"`
	PickupCommunityArea      int     `json:"pickup_community_area"`
	DropoffCommunityArea     int     `json:"dropoff_community_area"`
	Fare                     float64 `json:"fare"`
	Tip                      float64 `json:"tip"`
	AdditionalCharges        float64 `json:"additional_charges"`
	TripTotal                float64 `json:"trip_total"`
	SharedTripAuthorized     bool    `json:"shared_trip_authorized"`
	TripsPooled              int     `json:"trips_pooled"`
	PickupCentroidLatitude   float64 `json:"pickup_centroid_latitude"`
	PickupCentroidLongitude  float64 `json:"pickup_centroid_longitude"`
	PickupCentroidLocation   string  `json:"pickup_centroid_location"`
	DropoffCentroidLatitude  float64 `json:"dropoff_centroid_latitude"`
	DropoffCentroidLongitude float64 `json:"dropoff_centroid_longitude"`
	DropoffCentroidLocation  string  `json:"dropoff_centroid_location"`
}

// buildingPermits
func getBuildingPermitsData() {
	const (
		url    = "https://data.cityofchicago.org/resource/ydr8-5enu.json?$limit=500"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
	// Set up database connection
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Make API request to City of Chicago Building Permits dataset
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse JSON response into slice of buildingPermit structs
	var buildingPermits []buildingPermit
	err = json.NewDecoder(resp.Body).Decode(&buildingPermits)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Insert each building permit into the database
	for _, permit := range buildingPermits {
		query := `INSERT INTO buildingPermit (ID, PermitNumber, PermitType, ReviewType, TotalFee, AppStartDate, IssueDate, CommunityArea, Latitude, Longitude)
				  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

		_, err := db.Exec(query, permit.ID, permit.PermitNumber, permit.PermitType, permit.ReviewType, permit.TotalFee, permit.AppStartDate, permit.IssueDate, permit.CommunityArea, permit.Latitude, permit.Longitude)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Building permits data has been successfully inserted into the database.")
}

func getCCVIdata() {
	const (
		url    = "https://data.cityofchicago.org/resource/xhc6-88s9.json"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
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
		query := `INSERT INTO ccvi (GeographyType, CommunityAreaOrZip, CommunityAreaName, CcviCategory, CcviScore,
					SocioeconomicStatus, HouseholdComposition, NoPrimaryCare, CumMobilityRatio,
					FrontlineWorkers, Age65OrGreater, ComorbidConditions, CovidIncidenceRate,
					CovidHospitalRate, CrudeMortalityRate)
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

// censusData
func getCensusData() {
	const (
		url    = "https://data.cityofchicago.org/resource/kn9c-c2s2.json"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
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
		query := `INSERT INTO censusData (CommunityAreaNumber, CommunityAreaName, PercentOfHousingCrowded, PercentHouseholdsBelowPoverty,
			PercentAged16Unemployed, PercentAged25WithoutHighSchoolDiploma, PercentAgedUnder18OrOver64,
			PerCapitaIncome, HardshipIndex)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
		_, err := db.Exec(query, d.CommunityAreaNumber, d.CommunityAreaName, d.PercentOfHousingCrowded, d.PercentHouseholdsBelowPoverty,
			d.PercentAged16Unemployed, d.PercentAged25WithoutHighSchoolDiploma, d.PercentAgedUnder18OrOver64,
			d.PerCapitaIncome, d.HardshipIndex)
		if err != nil {
			panic(err)
		}
	}
}

// // Community Bound
func getCommunityBoundData() {
	// Open the shapefile
	file, err := os.Open("Boundaries_Community_Areas_current.shp")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the shapefile data
	shape, err := shp.NewDecoder(file).DecodeAll()
	if err != nil {
		panic(err)
	}

	// Connect to the Postgres database
	db, err := sql.Open("postgres", "user=postgres password=root dbname=chicagobi sslmode=require")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert each shape into the Postgres database
	for _, feature := range shape.Features {
		// Check if all columns are not blank
		if feature.Property[0].Value.String() == "" || feature.Property[1].Value.String() == "" {
			continue
		}

		// Insert the row into the Postgres database
		_, err := db.Exec("INSERT INTO communityBound (AreaNumber, AreaName, Geometry) VALUES ($1, $2, ST_GeomFromText($3, 4326))",
			feature.Property[0].Value, feature.Property[1].Value, feature.Geometry)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Successfully inserted data into Postgres database.")
}

// covid Cases
func getcovidCasesZipData() {
	const (
		url    = "https://data.cityofchicago.org/resource/yhhz-zm2v.json?$limit=500"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
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
	for _, row := range rows {
		_, err := db.Exec(`
			INSERT INTO covidCasesZip (
				Zip, 
				WeekNumber, 
				WeekStart, 
				WeekEnd, 
				CasesWeekly, 
				CasesMonthly, 
				CasesCum, 
				CaseRateWkly, 
				CaseRateMnth, 
				CaseRateCum, 
				TestsWeekly, 
				TestsCum, 
				TestRateWkly, 
				TestRateCum, 
				PosRateWkly, 
				PosRateCum, 
				DeathsWkly, 
				DeathsCum, 
				DeathRtWkly, 
				DeathRtCum, 
				Population, 
				RowId
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		`, row.Zip, row.WeekNumber, row.WeekStart, row.WeekEnd, row.CasesWeekly, row.CasesMonthly, row.CasesCum, row.CaseRateWkly, row.CaseRateMnth, row.CaseRateCum, row.TestsWeekly, row.TestsCum, row.TestRateWkly, row.TestRateCum, row.PosRateWkly, row.PosRateCum, row.DeathsWkly, row.DeathsCum, row.DeathRtWkly, row.DeathRtCum, row.Population, row.RowId)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Data has been successfully inserted into the database.")
}

// pubHealthStats
func getPublicHealthData() {
	const (
		url    = "https://data.cityofchicago.org/resource/iqnk-2tcu.json?$limit=500"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
	// Connect to the database
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Get data from API
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var rows []PublicHealthStats
	err = json.Unmarshal(body, &rows)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		_, err := db.Exec("INSERT INTO pubHealthStats (CommunityArea, CommunityAreaName, BirthRate, GeneralFertilityRate, LowBirthWeight, PrenatalCareBeginningInFirstTrimester, PretermBirths, TeenBirthRate, AssaultHomicide, BreastCancerInFemales, CancerAllSites, ColorectalCancer, DiabetesRelated, FirearmRelated, InfantMortalityRate, LungCancer, ProstateCancerInMales, StrokeCerebrovascularDisease, ChildhoodBloodLeadLevelScreening, ChildhoodLeadPoisoning, GonorrheaInFemales, GonorrheaInMales, Tuberculosis, BelowPovertyLevel, CrowdedHousing, Dependency, NoHighSchoolDiploma, PerCapitaIncome, Unemployment) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)",
			row.CommunityArea, row.CommunityAreaName, row.BirthRate, row.GeneralFertilityRate, row.LowBirthWeight, row.PrenatalCareBeginningInFirstTrimester, row.PretermBirths, row.TeenBirthRate, row.AssaultHomicide, row.BreastCancerInFemales, row.CancerAllSites, row.ColorectalCancer, row.DiabetesRelated, row.FirearmRelated, row.InfantMortalityRate, row.LungCancer, row.ProstateCancerInMales, row.StrokeCerebrovascularDisease, row.ChildhoodBloodLeadLevelScreening, row.ChildhoodLeadPoisoning, row.GonorrheaInFemales, row.GonorrheaInMales, row.Tuberculosis, row.BelowPovertyLevel, row.CrowdedHousing, row.Dependency, row.NoHighSchoolDiploma, row.PerCapitaIncome, row.Unemployment)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Public Health data has been successfully inserted into the database.")
}

// Taxi Trips
func getTaxiTripsData() {
	const (
		url    = "https://data.cityofchicago.org/resource/wrvz-psew.json?$limit=500"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
	// Connect to the database
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Get data from API
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var rows []TaxiTrips
	err = json.Unmarshal(body, &rows)
	if err != nil {
		panic(err)
	}

	// Iterate over the TaxiTrip slice and insert each row into the database
	for _, row := range rows {
		_, err := db.Exec(
			"INSERT INTO taxiTrips (TripID, TaxiID, TripStartTimestamp, TripEndTimestamp, TripSeconds, TripMiles, PickupCensusTract, DropoffCensusTract, PickupCommunityArea, DropoffCommunityArea, Fare, Tips, Tolls, Extras, TripTotal, PaymentType, Company, PickupCentroidLatitude, PickupCentroidLongitude, PickupCentroidLocation, DropoffCentroidLatitude, DropoffCentroidLongitude, DropoffCentroidLocation) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)",
			row.TripID, row.TaxiID, row.TripStartTimestamp, row.TripEndTimestamp, row.TripSeconds, row.TripMiles, row.PickupCensusTract, row.DropoffCensusTract, row.PickupCommunityArea, row.DropoffCommunityArea, row.Fare, row.Tips, row.Tolls, row.Extras, row.TripTotal, row.PaymentType, row.Company, row.PickupCentroidLatitude, row.PickupCentroidLongitude, row.PickupCentroidLocation, row.DropoffCentroidLatitude, row.DropoffCentroidLongitude, row.DropoffCentroidLocation,
		)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Inserted", len(rows), "rows into database")
}

// transNet
func getTransNetTripsData() {
	const (
		url    = "https://data.cityofchicago.org/resource/m6dm-c72p.json?$limit=500"
		dbHost = "localhost"
		dbPort = 5432
		dbUser = "postgres"
		dbPass = "root"
		dbName = "chicagobi"
	)
	// Connect to the database
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Get data from API
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var rows []TNPTrips
	err = json.Unmarshal(body, &rows)
	if err != nil {
		panic(err)
	}
	for _, row := range rows {
		_, err := db.Exec(`
			INSERT INTO transNetProviderTrips (
			TripID, 
			TripStartTimestamp, 
			TripEndTimestamp, 
			TripSeconds, 
			TripMiles, 
			PickupCensusTract, 
			DropoffCensusTract, 
			PickupCommunityArea, 
			DropoffCommunityArea, 
			Fare, 
			Tip, 
			AdditionalCharges, 
			TripTotal, 
			SharedTripAuthorized, 
			TripsPooled, 
			PickupCentroidLatitude, 
			PickupCentroidLongitude, 
			PickupCentroidLocation, 
			DropoffCentroidLatitude, 
			DropoffCentroidLongitude, 
			DropoffCentroidLocation)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`,
			row.TripID, row.TripStartTimestamp, row.TripEndTimestamp, row.TripSeconds, row.TripMiles,
			row.PickupCensusTract, row.DropoffCensusTract, row.PickupCommunityArea, row.DropoffCommunityArea, row.Fare,
			row.Tip, row.AdditionalCharges, row.TripTotal, row.SharedTripAuthorized, row.TripsPooled,
			row.PickupCentroidLatitude, row.PickupCentroidLongitude, row.PickupCentroidLocation,
			row.DropoffCentroidLatitude, row.DropoffCentroidLongitude, row.DropoffCentroidLocation)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Inserted", len(rows), "rows into database")
}

func main() {
	getBuildingPermitsData()
	//getData.getCCVIdata()
	//getData.getCensusData()
	//getData.getCommunityBoundData()
	//getData.getcovidCasesZipData()
	//getData.getPublicHealthData()
	//getData.getTaxiTripsData()
	//getData.getTransNetTripsData()
	// set time
	//s.Every(1).Day().At("00:05").Do(func(){ ... })

}
