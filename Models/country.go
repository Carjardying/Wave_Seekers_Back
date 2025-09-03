package Models

import (
	"context"
	"database/sql"
)

func AddCountry(db *sql.DB, c *Country) (int64, error) {
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
