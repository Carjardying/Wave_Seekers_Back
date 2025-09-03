package Models

import (
	"context"
	"database/sql"
)

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
