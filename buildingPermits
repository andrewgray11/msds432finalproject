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
	PermitSubType       string `json:"permit_subtype"`
	PermitDescription   string `json:"permit_description"`
	PermitIssueDate     string `json:"issue_date"`
	PermitEstimatedCost int    `json:"estimated_cost"`
	PermitStatus        string `json:"status"`
	PermitStreetNumber  int    `json:"street_number"`
	PermitStreetName    string `json:"street_direction"`
	PermitSuffix        string `json:"suffix"`
	PermitWorkType      string `json:"work_type"`
	PermitPIN1          int    `json:"pin1"`
	PermitPIN2          int    `json:"pin2"`
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
		// Check if all columns are not blank
		if permit.PermitNumber != "" && permit.PermitType != "" && permit.PermitSubType != "" && permit.PermitIssueDate != "" {
			_, err = db.Exec("INSERT INTO building_permits (id, permit_number, permit_type, permit_subtype, permit_description, issue_date, estimated_cost, status, street_number, street_name, suffix, work_type, pin1, pin2) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", permit.ID, permit.PermitNumber, permit.PermitType, permit.PermitSubType, permit.PermitDescription, permit.PermitIssueDate, permit.PermitEstimatedCost, permit.PermitStatus, permit.PermitStreetNumber, permit.PermitStreetName, permit.PermitSuffix, permit.PermitWorkType, permit.PermitPIN1, permit.PermitPIN2)
			if err != nil {
				fmt.Println("Error inserting row into database:", err)
				return
			}
		}
	}

	fmt.Println("Inserted", len(buildingPermits), "rows into database")
}
