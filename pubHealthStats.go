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
	url    = "https://data.cityofchicago.org/resource/iqnk-2tcu.json?$limit=500"
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "root"
	dbName = "master"
)

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

func main() {
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
