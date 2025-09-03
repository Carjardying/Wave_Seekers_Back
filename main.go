package main

import (
	"log"

	"example/Wave_Seekers_Back/Models"
)

func main() {
	dbFile := "waveseekers-database.db"

	// Crée le fichier de base de données s'il n'existe pas
	CreateDatabase(dbFile)

	// Connexion à la base
	db, err := Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Création des tables
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

	// Seed Users
	if err := Models.SeedUsers(db); err != nil {
		log.Fatal(err)
	}

	// Seed Countries
	if err := Models.SeedCountries(db); err != nil {
		log.Fatal(err)
	}

	// Seed Spots
	if err := Models.SeedSpots(db); err != nil {
		log.Fatal(err)
	}
}