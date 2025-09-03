package main

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

// CreateDatabase crée le fichier SQLite s'il n'existe pas
func CreateDatabase(dbFile string) {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Printf("Creating %s...\n", dbFile)
		file, err := os.Create(dbFile)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		log.Printf("%s created\n", dbFile)
	}
}

// Connect ouvre la connexion à la base SQLite
func Connect(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:"+dbFile+"?_busy_timeout=5000&_fk=1")
	if err != nil {
		return nil, err
	}
	return db, nil
}
