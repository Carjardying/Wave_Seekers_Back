package Models

import (
	"context"
	"database/sql"
	"log"
)

type Spot struct {
	ID              int     `json:"id" form:"id"`
	UserID          int     `json:"user_id" form:"user_id"`
	CountryID       int     `json:"country_id" form:"country_id"`
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

func GetAllSpots(db *sql.DB) ([]Spot, error) {
	rows, err := db.Query(`SELECT id, user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url FROM spot`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []Spot
	for rows.Next() {
		var spot Spot
		err := rows.Scan(&spot.ID, &spot.UserID, &spot.CountryID, &spot.Destination, &spot.Location, &spot.Lat, &spot.Long, &spot.PeakSeasonStart, &spot.PeakSeasonEnd, &spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL)
		if err != nil {
			return nil, err
		}
		spots = append(spots, spot)
	}
	return spots, nil
}

func GetSpotByID(db *sql.DB, id int) (*Spot, error) {
	spot := &Spot{}
	err := db.QueryRow(`SELECT id, user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url FROM spot WHERE id = ?`, id).Scan(&spot.ID, &spot.UserID, &spot.CountryID, &spot.Destination, &spot.Location, &spot.Lat, &spot.Long, &spot.PeakSeasonStart, &spot.PeakSeasonEnd, &spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL)
	if err != nil {
		return nil, err
	}
	return spot, nil
}

func GetSpotsByCountryID(db *sql.DB, countryID int) ([]Spot, error) {
	rows, err := db.Query(`SELECT id, user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url FROM spot WHERE country_id = ?`, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []Spot
	for rows.Next() {
		var spot Spot
		err := rows.Scan(&spot.ID, &spot.UserID, &spot.CountryID, &spot.Destination, &spot.Location, &spot.Lat, &spot.Long, &spot.PeakSeasonStart, &spot.PeakSeasonEnd, &spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL)
		if err != nil {
			return nil, err
		}
		spots = append(spots, spot)
	}
	return spots, nil
}
