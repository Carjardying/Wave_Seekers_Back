package Controllers

import (
	"context"

	"database/sql"

	"strings"

	"errors"

	"golang.org/x/crypto/bcrypt"

	"example/Wave_Seekers_Back/Models"
	token "example/Wave_Seekers_Back/Utils/Token"
)

func GetCurrentUserByID(db *sql.DB, uid uint) (Models.User, error) {
	var u Models.User

	err := db.QueryRow(`SELECT id, email, password FROM user WHERE id = ?`, uid).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, errors.New("user not found")
		}
		return u, err
	}

	u.PrepareGive()
	return u, nil
}

/*-------------------POST-------------------*/

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(db *sql.DB, email string, password string) (string, error) {

	u := Models.User{}

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

func AddUser(db *sql.DB, u *Models.User) (int64, error) {
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

/*-------------------GET-------------------*/

func GetAllUsers(db *sql.DB) ([]Models.User, error) {
	rows, err := db.Query(`SELECT id, email, password FROM user`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Models.User
	for rows.Next() {
		var user Models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(db *sql.DB, id int) (*Models.User, error) {
	user := &Models.User{}
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
