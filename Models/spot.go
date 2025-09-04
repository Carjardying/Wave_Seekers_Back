package Models

import (
	"context"
	"database/sql"
	"log"
)

type Spot struct {
	UserID          int
	CountryID       int
	Destination     string
	Location        string
	Lat             float64
	Long            float64
	PeakSeasonStart string
	PeakSeasonEnd   string
	DifficultyLevel int
	SurfingCulture  string
	ImageURL        string
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

func AddSpot(db *sql.DB, s *Spot) (int64, error) {
	var existingID int64
	err := db.QueryRow(`SELECT id FROM spot WHERE destination = ?`, s.Destination).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if existingID != 0 {
		return existingID, nil
	}
	result, err := db.ExecContext(
		context.Background(),
		`INSERT INTO spot (user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		s.UserID, s.CountryID, s.Destination, s.Location, s.Lat, s.Long, s.PeakSeasonStart, s.PeakSeasonEnd, s.DifficultyLevel, s.SurfingCulture, s.ImageURL,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
