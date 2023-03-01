package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

// Define a struct to hold the transportation network provider trip data
type TNPTrip struct {
	ID                         int
	TripID                     string
	PickupCommunityArea        string
	DropoffCommunityArea       string
	TripStartTimestamp         time.Time
	TripEndTimestamp           time.Time
	TripSeconds                int
	TripMiles                  float64
	PickupCensusTract          string
	DropoffCensusTract         string
	PickupCentroidLatitude     float64
	PickupCentroidLongitude    float64
	DropoffCentroidLatitude    float64
	DropoffCentroidLongitude   float64
	SharedTripAuthorized       string
	TripsPooled                int
	PickupCommunityName        string
	DropoffCommunityName       string
	Fare                       float64
	Tip                        float64
	AdditionalCharges          float64
	TripTotal                  float64
	SharedTripPaymentType      string
	PaymentType                string
	Company                    string
	TripType                   string
	TripExtras                 float64
	TripStartTimeAdjusted      time.Time
	TripEndTimeAdjusted        time.Time
	TripMilesAdjusted          float64
	TripFareAdjusted           float64
	TripTotalAdjusted          float64
	SharedTripActualDistance   float64
	SharedTripMatchedID        string
	OriginationLatitude        float64
	OriginationLongitude       float64
	DestinationLatitude        float64
	DestinationLongitude       float64
	SharedTripCost             float64
	NumberOfMatchedSharedTrips int
	// Add additional fields as needed
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

	// Make an API request to retrieve the transportation network provider trip data
	resp, err := client.Get("https://data.cityofchicago.org/resource/m6dm-c72p.json?$limit=500")
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of TNPTrip structs
	var tnpTrips []TNPTrip
	err = json.NewDecoder(resp.Body).Decode(&tnpTrips)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	// Iterate over the TNPTrip slice and insert each row into the database
	for _, trip := range tnpTrips {
		if trip.PickupCommunityArea == "" || trip.DropoffCommunityArea == "" ||
			trip.TripStartTimestamp.IsZero() || trip.TripEndTimestamp.IsZero() ||
			trip.TripSeconds == 0 || trip.TripMiles == 0 ||
			trip.Fare == 0 || trip.TripTotal == 0 {
			continue // Skip rows with blank columns
		}
		_, err = db.Exec("INSERT INTO tnp_trips (trip_id, pickup_community_area, dropoff_community_area, trip_start_timestamp, trip_end_timestamp, trip_seconds, trip_miles, pickup_census_tract, dropoff_census_tract, pickup_centroid_latitude, pickup_centroid_longitude, dropoff_centroid_latitude, dropoff_centroid_longitude, shared_trip_authorized, trips_pooled, pickup_community_name, dropoff_community_name, fare, tip, additional_charges, trip_total, shared_trip_payment_type, payment_type, company, trip_type, trip_extras, trip_start_time_adjusted, trip_end_time_adjusted, trip_miles_adjusted, trip_fare_adjusted, trip_total_adjusted, shared_trip_actual_distance, shared_trip_matched_id, origination_latitude, origination_longitude, destination_latitude, destination_longitude, shared_trip_cost, number_of_matched_shared_trips) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41)", trip.TripID, trip.PickupCommunityArea, trip.DropoffCommunityArea, trip.TripStartTimestamp, trip.TripEndTimestamp, trip.TripSeconds, trip.TripMiles, trip.PickupCensusTract, trip.DropoffCensusTract, trip.PickupCentroidLatitude, trip.PickupCentroidLongitude, trip.DropoffCentroidLatitude, trip.DropoffCentroidLongitude, trip.SharedTripAuthorized, trip.TripsPooled, trip.PickupCommunityName, trip.DropoffCommunityName, trip.Fare, trip.Tip, trip.AdditionalCharges, trip.TripTotal, trip.SharedTripPaymentType, trip.PaymentType, trip.Company, trip.TripType, trip.TripExtras, trip.TripStartTimeAdjusted, trip.TripEndTimeAdjusted, trip.TripMilesAdjusted, trip.TripFareAdjusted, trip.TripTotalAdjusted, trip.SharedTripActualDistance, trip.SharedTripMatchedID, trip.OriginationLatitude, trip.OriginationLongitude, trip.DestinationLatitude, trip.DestinationLongitude, trip.SharedTripCost, trip.NumberOfMatchedSharedTrips)

		// Handle any errors that occurred during the INSERT statement
		if err != nil {
			fmt.Println("Error inserting row into database:", err)
			return
		}
	}

	fmt.Println("Inserted", len(tnpTrips), "rows into database")
}
