package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// Define a struct to hold the taxi trip data
type TaxiTrip struct {
	ID                   int
	TripStartTimestamp   time.Time
	TripEndTimestamp     time.Time
	TripSeconds          int
	TripMiles            float64
	PickupCensusTract    string
	DropoffCensusTract   string
	PickupCommunityArea  string
	DropoffCommunityArea string
	Fare                 float64
	Tips                 float64
	Tolls                float64
	Extras               float64
	TripTotal            float64
	PaymentType          string
	Company              string
	TaxiID               string
}

func main() {
	// Set up a database connection
	db, err := sql.Open("postgres", "host=localhost dbname=my_database user=my_username password=my_password sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Set up an HTTP client to make API requests
	client := http.Client{
		Timeout: time.Second * 10,
	}

	// Make an API request to retrieve the taxi trip data
	resp, err := client.Get("https://data.cityofchicago.org/resource/wrvz-psew.json?$limit=500")
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of TaxiTrip structs
	var taxiTrips []TaxiTrip
	err = json.NewDecoder(resp.Body).Decode(&taxiTrips)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Iterate over the TaxiTrip slice and insert each row into the database
	for _, trip := range taxiTrips {
		if trip.TripStartTimestamp.IsZero() || trip.TripEndTimestamp.IsZero() ||
			trip.TripSeconds == 0 || trip.TripMiles == 0 || trip.PickupCommunityArea == "" ||
			trip.DropoffCommunityArea == "" || trip.Fare == 0 || trip.TripTotal == 0 {
			continue // Skip rows with blank columns
		}

		_, err := db.Exec("INSERT INTO taxi_trips (trip_start_timestamp, trip_end_timestamp, trip_seconds, trip_miles, pickup_census_tract, dropoff_census_tract, pickup_community_area, dropoff_community_area, fare, tips, tolls, extras, trip_total, payment_type, company, taxi_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)",
			trip.TripStartTimestamp, trip.TripEndTimestamp, trip.TripSeconds, trip.TripMiles, trip.PickupCensusTract, trip.DropoffCensusTract, trip.PickupCommunityArea, trip.DropoffCommunityArea, trip.Fare, trip.Tips, trip.Tolls, trip.Extras, trip.TripTotal, trip.PaymentType, trip.Company, trip.TaxiID)

		// Handle any errors that occurred during the INSERT statement
		if err != nil {
			fmt.Println("Error inserting row into database:", err)
			return
		}
	}
	fmt.Println("Inserted", len(taxiTrips), "rows into database")
}
