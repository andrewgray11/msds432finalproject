package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/relvacode/iso8601"
	_ "github.com/lib/pq"
)

const (
	url    = "https://data.cityofchicago.org/resource/wrvz-psew.json?$limit=500"
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "root"
	dbName = "master"
)

// Define a struct to hold the taxi trip data
type TaxiTrips struct {
	TripID                   string  `json:"trip_id"`
	TaxiID                   string  `json:"taxi_id"`
	TripStartTimestamp       time.Time  `json:"trip_start_timestamp"`
	TripEndTimestamp         time.Time  `json:"trip_end_timestamp"`
	TripSeconds              int     `json:"trip_seconds"`
	TripMiles                float64 `json:"trip_miles"`
	PickupCensusTract        string  `json:"pickup_census_tract"`
	DropoffCensusTract       string  `json:"dropoff_census_tract"`
	PickupCommunityArea      int     `json:"pickup_community_area"`
	DropoffCommunityArea     int     `json:"dropoff_community_area"`
	Fare                     float64 `json:"fare"`
	Tips                     float64 `json:"tips"`
	Tolls                    float64 `json:"tolls"`
	Extras                   float64 `json:"extras"`
	TripTotal                float64 `json:"trip_total"`
	PaymentType              string  `json:"payment_type"`
	Company                  string  `json:"company"`
	PickupCentroidLatitude   float64 `json:"pickup_centroid_latitude"`
	PickupCentroidLongitude  float64 `json:"pickup_centroid_longitude"`
	PickupCentroidLocation   string  `json:"pickup_centroid_location"`
	DropoffCentroidLatitude  float64 `json:"dropoff_centroid_latitude"`
	DropoffCentroidLongitude float64 `json:"dropoff_centroid_longitude"`
	DropoffCentroidLocation  string  `json:"dropoff_centroid_location"`
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
