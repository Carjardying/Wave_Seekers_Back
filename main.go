package main

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"

	"log"

	"example/Wave_Seekers_Back/Models"

	"strconv"
)

var db *sql.DB

func main() {

	dbFile := "waveseekers-database.db"

	// DB creation
	CreateDatabase(dbFile)

	// DB connexion
	var err error
	db, err = Connect(dbFile)
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
	// if err := Models.CreateLikedSpotTable(db); err != nil {
	// 	log.Fatal(err)
	// }

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

	router := gin.Default()
	router.GET("/users/:id", getUserHandler)
	router.Run("localhost:8080")
}

// Handler function that calls GetUserByID
func getUserHandler(c *gin.Context) {
	idStr := c.Param("id")

	// Converting int to string (like a ParsInt)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid user ID"})
		return
	}

	user, err := Models.GetUserByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching user"})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}
