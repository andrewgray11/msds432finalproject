package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

type CovidZip struct {
	Zip        string `json:"zip_code"`
	Cases      int    `json:"cases"`
	Tests      int    `json:"tests"`
	Deaths     int    `json:"deaths"`
	PeopleTest int    `json:"people_tested"`
}

func main() {
	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/mydatabase?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Make HTTP request to API
	resp, err := http.Get("https://data.cityofchicago.org/resource/yhhz-zm2v.json?$limit=500")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON response into slice of CovidZip structs
	var data []CovidZip
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Insert data into database
	for _, row := range data {
		if row.Zip != "" && row.Cases != 0 && row.Tests != 0 && row.Deaths != 0 && row.PeopleTest != 0 {
			query := fmt.Sprintf("INSERT INTO covid_zip (zip_code, cases, tests, deaths, people_tested) VALUES ('%s', %d, %d, %d, %d)", strings.TrimSpace(row.Zip), row.Cases, row.Tests, row.Deaths, row.PeopleTest)
			_, err := db.Exec(query)
			if err != nil {
				panic(err)
			}
		}
	}
}
