package Models

import (
	"context"
	"database/sql"
	"log"
)

type User struct {
	Email    string
	Password string
}

func CreateUserTable(db *sql.DB) error {
	ddl := `CREATE TABLE IF NOT EXISTS user (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP	
    );`
	_, err := db.Exec(ddl)
	log.Println("User Table created")
	return err
}

func AddUser(db *sql.DB, u *User) (int64, error) {
	var existingID int64
	err := db.QueryRow(`SELECT id FROM user WHERE email = ?`, u.Email).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if existingID != 0 {
		return existingID, nil
	}
	result, err := db.ExecContext(context.Background(), `INSERT INTO user (email, password) VALUES (?, ?)`, u.Email, u.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
