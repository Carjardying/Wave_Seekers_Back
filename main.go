package main

import (
	"log"

	"example/Wave_Seekers_Back/Models"
)

func main() {
	dbFile := "waveseekers-database.db"

	// DB creation
	CreateDatabase(dbFile)

	// DB connexion
	db, err := Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
	if err := Models.CreateUserTable(db); err != nil {
		log.Fatal(err)
	}
	if err := Models.CreateCountryTable(db); err != nil {
		log.Fatal(err)
	}
	if err := Models.CreateSpotTable(db); err != nil {
		log.Fatal(err)
	}
	if err := Models.CreateLikedSpotTable(db); err != nil {
		log.Fatal(err)
	}

	// Seeders
	if err := Models.SeedUsers(db); err != nil {
		log.Fatal(err)
	}

	if err := Models.SeedCountries(db); err != nil {
		log.Fatal(err)
	}

	if err := Models.SeedSpots(db); err != nil {
		log.Fatal(err)
	}
}
