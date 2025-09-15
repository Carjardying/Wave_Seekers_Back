package Models

import (
	"database/sql"

	"log"

	"html"

	"strings"

	"golang.org/x/crypto/bcrypt"
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

func (u *User) PrepareGive() {
	u.Password = ""
}

/*-------------------POST-------------------*/

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
