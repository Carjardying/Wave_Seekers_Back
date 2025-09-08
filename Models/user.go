package Models

import (
	"context"

	"database/sql"

	"log"

	"html"

	"strings"

	"errors"

	"golang.org/x/crypto/bcrypt"

	token "example/Wave_Seekers_Back/Utils/Token"
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

func GetCurrentUserByID(db *sql.DB, uid uint) (User, error) {
	var u User

	err := db.QueryRow(`SELECT id, email, password FROM user WHERE id = ?`, uid).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.New("User not found!")
		}
		return u, err
	}

	u.PrepareGive()
	return u, nil
}

func (u *User) PrepareGive() {
	u.Password = ""
}

/*-------------------POST-------------------*/

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(db *sql.DB, email string, password string) (string, error) {

	u := User{}

	err := db.QueryRow(`SELECT id, email, password FROM user WHERE email = ?`, email).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil {
		return "", err
	}

	tokenString, err := token.GenerateToken(uint(u.ID))
	if err != nil {
		return "", err
	}

	return tokenString, nil

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

	if !strings.HasPrefix(u.Password, "$2") {
		if err := u.BeforeAddUser(); err != nil {
			return 0, err
		}
	}

	result, err := db.ExecContext(context.Background(), `INSERT INTO user (email, password) VALUES (?, ?)`, u.Email, u.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (u *User) BeforeAddUser() error {

	//Turns password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//Removes spaces in email
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}

/*-------------------GET-------------------*/

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

/*-------------------UPDATE-------------------*/

/*-------------------DELETE-------------------*/

func DeleteUser(db *sql.DB, id int) error {
	// check if user exists
	var exists int
	err := db.QueryRow(`SELECT id FROM user WHERE id = ?`, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows // no user
		}
		return err // other error in DB
	}

	// delete the user
	result, err := db.ExecContext(
		context.Background(),
		`DELETE FROM user WHERE id = ?`,
		id,
	)
	if err != nil {
		return err
	}

	//check if user is deleted by checking rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil // Success
}
