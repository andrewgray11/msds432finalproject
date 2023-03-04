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
	url    = "https://data.cityofchicago.org/resource/m6dm-c72p.json?$limit=500"
	dbHost = "localhost"
	dbPort = 5432
	dbUser = "postgres"
	dbPass = "root"
	dbName = "master"
)

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
