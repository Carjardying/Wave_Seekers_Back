package Models

import (
    "database/sql"
    "log"
)

// Tables
func CreateUserTable(db *sql.DB) error {
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

func CreateCountryTable(db *sql.DB) error {
    ddl := `CREATE TABLE IF NOT EXISTS country (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
        name TEXT NOT NULL
    );`
    _, err := db.Exec(ddl)
    log.Println("Country Table created")
    return err
}

func CreateSpotTable(db *sql.DB) error {
    ddl := `CREATE TABLE IF NOT EXISTS spot (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER,
        country_id INTEGER,
        destination TEXT NOT NULL,
        location TEXT NOT NULL,
        lat REAL NOT NULL,
        long REAL NOT NULL,
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

func CreateLikedSpotTable(db *sql.DB) error {
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
