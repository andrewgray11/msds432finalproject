package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jonas-p/go-shp"
	_ "github.com/lib/pq"
)

func main() {
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
	db, err := sql.Open("postgres", "user=postgres password=root dbname=master sslmode=require")
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
