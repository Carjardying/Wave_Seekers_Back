package main

import (
	"database/sql"

	"net/http"

	"log"

	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"example/Wave_Seekers_Back/Models"

	"example/Wave_Seekers_Back/Seeders"

	"example/Wave_Seekers_Back/Controllers"
)

var db *sql.DB

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

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

	// Initialize Controllers with database connection
	Controllers.InitializeDB(db)

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
	if err := Seeders.SeedUsers(db); err != nil {
		log.Fatal(err)
	}

	if err := Seeders.SeedCountries(db); err != nil {
		log.Fatal(err)
	}

	if err := Seeders.SeedSpots(db); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/users/:id", getUserByIDHandler)
	router.GET("/spots", getAllSpotsHandler)
	router.GET("/spots/:id", getSpotByIDHandler) //Spot's Details
	router.GET("/spots/country/:country_id", getSpotByCountryHandler)
	router.GET("/spots/user/:user_id", getSpotsByUserIDHandler)

	router.POST("/spots", addSpotHandler)
	router.POST("/signup", Controllers.SignUp)
	router.POST("/login", Controllers.Login)

	router.DELETE("/users/:user_id", deleteUserHandler)

	router.Run("localhost:8080")
}

/*---------- GET------*/

// Handler function that calls GetUserByID
func getUserByIDHandler(c *gin.Context) {
	idStr := c.Param("id")

	// Converting int to string for id(like a ParsInt)
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

// GetAllSpot's Handler
func getAllSpotsHandler(c *gin.Context) {

	spot, err := Models.GetAllSpots(db)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "spots list not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching spots list"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, spot)

}

func getSpotByIDHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid spot ID"})
		return
	}

	spot, err := Models.GetSpotByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "spot not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching spot"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, spot)
}

func getSpotByCountryHandler(c *gin.Context) {
	idStr := c.Param("country_id")

	country_id, err := strconv.Atoi(idStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid spot ID"})
		return
	}

	spot, err := Models.GetSpotsByCountryID(db, country_id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "spot not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching spot"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, spot)
}

func getSpotsByUserIDHandler(c *gin.Context) {
	idStr := c.Param("user_id")

	user_id, err := strconv.Atoi(idStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid spot ID"})
		return
	}

	spot, err := Models.GetSpotsByUserID(db, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "spot not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching spot"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, spot)

}

/*---------- POST------*/

func addSpotHandler(c *gin.Context) {
	var spot Models.Spot

	if err := c.ShouldBind(&spot); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid form data", "details": err.Error()})
		return
	}

	if spot.Destination == "" || spot.Location == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Destination and Location are required"})
		return
	}

	id, err := Models.AddSpot(db, &spot)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create spot", "details": err.Error()})
		return
	}

	spot.ID = int(id)
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Spot created successfully",
		"spot":    spot,
		"id":      id,
	})
}

/*---------------DELETE-------------*/

func deleteUserHandler(c *gin.Context) {
	idStr := c.Param("user_id")

	user_id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid User ID"})
		return
	}

	err = Models.DeleteUser(db, user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error deleting user", "error": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message":    "user deleted successfully",
		"deleted_id": user_id,
	})
}
