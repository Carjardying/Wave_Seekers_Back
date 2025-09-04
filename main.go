package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"example/Wave_Seekers_Back/Models"
)

var db *sql.DB

func main() {
    router := gin.Default()
    router.GET("/users/1", getUserInfo)

    router.Run("localhost:8080")

	dbFile := "waveseekers-database.db"

	// DB creation
	CreateDatabase(dbFile)

	// DB connexion
	db, err := Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
	if err := Models.CreateUserTable(db); err != nil {
		log.Fatal(err)
	}
	if err := Models.CreateCountryTable(db); err != nil {
		log.Fatal(err)
	}
	if err := Models.CreateSpotTable(db); err != nil {
		log.Fatal(err)
	}
	if err := Models.CreateLikedSpotTable(db); err != nil {
		log.Fatal(err)
	}

	// Seeders
	if err := Models.SeedUsers(db); err != nil {
		log.Fatal(err)
	}

	if err := Models.SeedCountries(db); err != nil {
		log.Fatal(err)
	}

	if err := Models.SeedSpots(db); err != nil {
		log.Fatal(err)
	}
}

// func getUserInfo(u *gin.Context) {
//       id := u.Param("id")

//     users, err := Models.SeedUsers(db *sql.DB) // pass db
//     if err != nil {
//         u.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching users"})
//         return
//     }

//     for _, a := range users {
//         if a.ID == id {
//             u.IndentedJSON(http.StatusOK, a)
//             return
//         }
//     }
//     u.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
// }



func getUserInfo(u *gin.Context)([]Models.User, error) {
    // c.JSON(200, gin.H{"message":"Get One User Info"})
    // rows, err := DB.Query(`SELECT * FROM user WHERE id = ?`)
    // if err !=nil {
    //     return nil, err
    // }
    // return Users

     id := u.Param("id")

    users, err := Models.User(db *sql.DB)
    if err != nil {
        u.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching users"})
        return nil, err
    }

    for _, a := range users {
        if a.ID == id {
            u.IndentedJSON(http.StatusOK, a)
            return users
        }
    }
    u.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
	
