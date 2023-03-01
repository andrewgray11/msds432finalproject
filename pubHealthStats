package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    _ "github.com/lib/pq"
)

type PublicHealthStats struct {
    CommunityArea      string `json:"community_area"`
    BirthRate          string `json:"birth_rate"`
    GeneralFertility   string `json:"general_fertility_rate"`
    LowBirthWeight     string `json:"low_birth_weight"`
    PrenatalCare       string `json:"prenatal_care_beginning_in_first_trimester"`
    TeenBirthRate      string `json:"teen_birth_rate"`
    Uninsured          string `json:"uninsured"`
    BelowPoverty       string `json:"below_poverty_level"`
    CrowdedHousing     string `json:"crowded_housing"`
    Dependency         string `json:"dependency"`
    NoDiploma          string `json:"no_high_school_diploma"`
    PerCapitaIncome    string `json:"per_capita_income"`
    Unemployment       string `json:"unemployment"`
    Assault            string `json:"assault"`
    BreastCancer       string `json:"breast_cancer_in_females"`
    Cancer             string `json:"cancer"`
    ColorectalCancer   string `json:"colorectal_cancer"`
    Diabetes           string `json:"diabetes_related"`
    FirearmMortality   string `json:"firearm_related"`
    InfantMortality    string `json:"infant_mortality_rate"`
    LungCancer         string `json:"lung_cancer"`
    ProstateCancer     string `json:"prostate_cancer_in_males"`
    Stroke             string `json:"stroke_cerebrovascular_disease"`
    Tuberculosis       string `json:"tuberculosis"`
    BelowPovertyRecent string `json:"below_poverty_level_recent"`
    NoDiplomaRecent    string `json:"no_high_school_diploma_recent"`
    UnemploymentRecent string `json:"unemployment_recent"`
}

func main() {
    // Connect to the database
    connStr := "postgres://username:password@localhost/database?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Set up the HTTP client and request
    client := http.Client{}
    req, err := http.NewRequest("GET", "https://data.cityofchicago.org/resource/iqnk-2tcu.json?$limit=500", nil)
    if err != nil {
        log.Fatal(err)
    }

    // Add the app token to the request header
    req.Header.Add("X-App-Token", "YOUR_APP_TOKEN_HERE")

    // Send the request and parse the response
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    var stats []PublicHealthStats
    err = json.NewDecoder(resp.Body).Decode(&stats)
    if err != nil {
        log.Fatal(err)
    }

    for _, record := range data {
        // Check if all columns are not blank
        if record.CommunityArea != "" && record.CancerCrudeRate != "" && record.InfantMortalityRate != "" && record.TeenBirthRate != "" && record.UnemploymentRate != "" && record.LowBirthWeight != "" && record.Smoker != "" {
            // Insert record into the database
            err = insertRecord(db, record)
            if err != nil {
                log.Printf("Error inserting record: %v", err)
            }
        }
    }
        
