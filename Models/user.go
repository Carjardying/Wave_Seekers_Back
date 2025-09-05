package Models

import (
	"context"
	"database/sql"
	"log"
)

type User struct {
	ID       int
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

func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, email, password FROM user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	user := &User{}
	err := db.QueryRow(`SELECT id, email, password FROM user WHERE id = ?`, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
