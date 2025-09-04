package Models

import (
	"database/sql"
	"log"
)

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
