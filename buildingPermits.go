package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type buildingPermit struct {
	ID                  int    `json:"id"`
	PermitNumber        string `json:"permit_"` // use backtick to handle column name with underscore
	PermitType          string `json:"permit_type"`
	ReviewType			string `json:"review_type"`
	TotalFee			float64`json:"total_fee"`
	AppStartDate		string `json:"application_start_date"`
	IssueDate			string `json:"issue_date"`
	CommunityArea		int	   `json:"community_area"`
	Latitude			float64`json:"latitude"`
	Longitude			float64`json:"longitude"`
}

func main() {
	// Set up database connection
	db, err := sql.Open("postgres", "postgres://user:password@host/database?sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	// Make API request to City of Chicago Building Permits dataset
	resp, err := http.Get("https://data.cityofchicago.org/resource/ydr8-5enu.json?$limit=500")
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
		query := `INSERT INTO building_permits (id, permit_number, permit_type, review_type, total_fee, application_start_date, issue_date, community_area, latitude, longitude)
				  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	
		_, err := db.Exec(query, permit.ID, permit.PermitNumber, permit.PermitType, permit.ReviewType, permit.TotalFee, permit.AppStartDate, permit.IssueDate, permit.CommunityArea, permit.Latitude, permit.Longitude)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Building permits data has been successfully inserted into the database.")
}
