package main

import (
	"database/sql"
	"log"
	"os"

    _ "github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
    db, err := sql.Open("sqlite", "file:waveseekers-database.db?_busy_timeout=5000&_fk=1") // Open the created SQLite File
	if err != nil {
        panic(err)
    }
    
    defer db.Close() // Defer Closing the database
	if err := createUserTable(db); err != nil {
        panic(err)
    } // Create Database Tables
}

func createDatabase() {
	log.Println("Creating waveseekers-database.db...")
	file, err := os.Create("waveseekers-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("waveseekers-database.db created")

	// insertUser(db, "leoniemiege@gmail.com", "test")
}

func createUserTable(db *sql.DB) error {
    ddl := `CREATE TABLE IF NOT EXISTS user (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP	
    );`
    
    _, err := db.Exec(ddl)
    return err
}

// func insertUser(db *sql.DB, email string, password string) {
// 	log.Println("Inserting user record ...")
// 	insertUserSQL := `INSERT INTO user(email, password) VALUES (?, ?)`
// 	statement, err := db.Prepare(insertUserSQL) // Prepare statement. 
//                                                    // This is good to avoid SQL injections
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	_, err = statement.Exec(email, password)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// }
