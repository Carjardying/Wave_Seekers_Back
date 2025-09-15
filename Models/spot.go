package Models

import (
	"database/sql"
	"log"
)

type Spot struct {
	ID              int     `json:"id" form:"id"`
	UserID          int     `json:"user_id" form:"user_id"`
	CountryID       int     `json:"country_id" form:"country_id"`
	CountryName     string  `json:"country_name"`
	Destination     string  `json:"destination" form:"destination"`
	Location        string  `json:"location" form:"location"`
	Lat             float64 `json:"lat" form:"lat"`
	Long            float64 `json:"long" form:"long"`
	PeakSeasonStart string  `json:"peak_season_start" form:"peak_season_start"`
	PeakSeasonEnd   string  `json:"peak_season_end" form:"peak_season_end"`
	DifficultyLevel int     `json:"difficulty_level" form:"difficulty_level"`
	SurfingCulture  string  `json:"surfing_culture" form:"surfing_culture"`
	ImageURL        string  `json:"image_url" form:"image_url"`
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
