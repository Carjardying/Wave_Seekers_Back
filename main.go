package main

import (
	"database/sql"
	"log"
	"os"

    _ "github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

type User struct {
	Email string
	Password string
}

func main() {
    db, err := sql.Open("sqlite", "file:waveseekers-database.db?_busy_timeout=5000&_fk=1") // Open the created SQLite File
	if err != nil {
        panic(err)
    }
    
    defer db.Close() // Defer Closing the database
	if err := createUserTable(db); err != nil {
        panic(err)
    } // Create UserTable

    if err := createCountryTable(db); err != nil {
        panic(err)
    } // Create CountryTable

    if err := createSpotTable(db); err != nil {
        panic(err)
    } // Create SpotTable

    if err := createLikedSpotTable(db); err != nil {
        panic(err)
    } // Create LikedSpotTable
}

func createDatabase() {
	log.Println("Creating waveseekers-database.db...")
	file, err := os.Create("waveseekers-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("waveseekers-database.db created")
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
    log.Println("User Table created")
    return err
}

func createCountryTable(db *sql.DB) error {
    ddl := `CREATE TABLE IF NOT EXISTS country (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		name TEXT NOT NULL
    );`
    
    _, err := db.Exec(ddl)
    log.Println("Country Table created")
    return err
}

func createSpotTable(db *sql.DB) error {
    ddl := `CREATE TABLE IF NOT EXISTS spot (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        country_id INTEGER,
        destination TEXT NOT NULL,
        location TEXT NOT NULL,
        long TEXT NOT NULL,
        lat TEXT NOT NULL,
        peak_season_start TEXT NOT NULL,
        peak_season_end TEXT NOT NULL,
        difficulty_level INTEGER NOT NULL,
        surfing_culture TEXT NOT NULL,
        image_url TEXT NOT NULL,		
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES user(id),
        FOREIGN KEY (country_id) REFERENCES country(id)
    );`
    
    _, err := db.Exec(ddl)
    log.Println("Spot Table created")
    return err
}

func createLikedSpotTable(db *sql.DB) error {
    ddl := `CREATE TABLE IF NOT EXISTS liked_spot (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		user_id INTEGER,
        spot_id INTEGER,
        FOREIGN KEY (user_id) REFERENCES user(id),
        FOREIGN KEY (spot_id) REFERENCES spot(id)
    );`
    
    _, err := db.Exec(ddl)
    log.Println("Liked Spot Table created")
    return err
}

func addUser(u *User) (int64, error) {
	result, err := db.ExecContext(
		context.Background(),
		`INSERT INTO user (email, password) VALUES (?,?);`, u.Email, u.Password
	)
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

