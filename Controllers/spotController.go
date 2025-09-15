package Controllers

import (
	"context"
	"database/sql"
	"example/Wave_Seekers_Back/Models"
)

/*-------------------POST-------------------*/

func AddSpot(db *sql.DB, s *Models.Spot) (int64, error) {
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

/*-------------------GET-------------------*/

func GetAllSpots(db *sql.DB) ([]Models.Spot, error) {
	rows, err := db.Query(`
        SELECT 
            s.id, s.user_id, s.country_id,
            s.destination, s.location, s.lat, s.long,
            s.peak_season_start, s.peak_season_end,
            s.difficulty_level, s.surfing_culture, s.image_url,
            COALESCE(c.name, s.location, 'Unknown') AS country_name
        FROM spot s
        LEFT JOIN country c ON s.country_id = c.id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []Models.Spot
	for rows.Next() {
		var spot Models.Spot
		var countryName sql.NullString // <-- important

		if err := rows.Scan(
			&spot.ID, &spot.UserID, &spot.CountryID,
			&spot.Destination, &spot.Location, &spot.Lat, &spot.Long,
			&spot.PeakSeasonStart, &spot.PeakSeasonEnd,
			&spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL,
			&countryName,
		); err != nil {
			return nil, err
		}

		if countryName.Valid && countryName.String != "" {
			spot.CountryName = countryName.String
		} else {
			spot.CountryName = "Unknown"
		}

		spots = append(spots, spot)
	}
	return spots, nil
}

func GetSpotByID(db *sql.DB, id int) (*Models.Spot, error) {
	spot := &Models.Spot{}
	err := db.QueryRow(`SELECT id, user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url FROM spot WHERE id = ?`, id).Scan(&spot.ID, &spot.UserID, &spot.CountryID, &spot.Destination, &spot.Location, &spot.Lat, &spot.Long, &spot.PeakSeasonStart, &spot.PeakSeasonEnd, &spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL)
	if err != nil {
		return nil, err
	}
	return spot, nil
}

func GetSpotsByCountryID(db *sql.DB, countryID int) ([]Models.Spot, error) {
	rows, err := db.Query(`SELECT id, user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url FROM spot WHERE country_id = ?`, countryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []Models.Spot
	for rows.Next() {
		var spot Models.Spot
		err := rows.Scan(&spot.ID, &spot.UserID, &spot.CountryID, &spot.Destination, &spot.Location, &spot.Lat, &spot.Long, &spot.PeakSeasonStart, &spot.PeakSeasonEnd, &spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL)
		if err != nil {
			return nil, err
		}
		spots = append(spots, spot)
	}
	return spots, nil
}

// Select spot by user id
func GetSpotsByUserID(db *sql.DB, userID int) ([]Models.Spot, error) {
	rows, err := db.Query(`SELECT id, user_id, country_id, destination, location, lat, long, peak_season_start, peak_season_end, difficulty_level, surfing_culture, image_url FROM spot WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []Models.Spot
	for rows.Next() {
		var spot Models.Spot
		err := rows.Scan(&spot.ID, &spot.UserID, &spot.CountryID, &spot.Destination, &spot.Location, &spot.Lat, &spot.Long, &spot.PeakSeasonStart, &spot.PeakSeasonEnd, &spot.DifficultyLevel, &spot.SurfingCulture, &spot.ImageURL)
		if err != nil {
			return nil, err
		}
		//la table spots = append(la table spots, le spot)
		spots = append(spots, spot)
	}
	return spots, nil
}

/*-------------------UPDATE-------------------*/

/*-------------------DELETE-------------------*/
