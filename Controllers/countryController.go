package Controllers

import (
	"context"
	"database/sql"
	"example/Wave_Seekers_Back/Models"
)

func AddCountry(db *sql.DB, c *Models.Country) (int64, error) {
	var existingID int64
	err := db.QueryRow(`SELECT id FROM country WHERE name = ?`, c.Name).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if existingID != 0 {
		return existingID, nil
	}
	result, err := db.ExecContext(context.Background(), `INSERT INTO country (name) VALUES (?)`, c.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetAllCountries(db *sql.DB) ([]Models.Country, error) {
	rows, err := db.Query(`SELECT id, name FROM country ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countries []Models.Country
	for rows.Next() {
		var country Models.Country
		err := rows.Scan(&country.ID, &country.Name)
		if err != nil {
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}
