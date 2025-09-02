package models

import (
	// "database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() {
	log.Println("Creating waveseekers-database.db...")
	file, err := os.Create("waveseekers-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("waveseekers-database.db created")
}
