package Controllers

import (
	"net/http"

	"database/sql"
  	
	"github.com/gin-gonic/gin"
	
	"example/Wave_Seekers_Back/Models"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
	db = database
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := Models.LoginCheck(db, input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is invalid."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

type SignUpInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context){
	
	var input SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := Models.User{}
	u.Email = input.Email
	u.Password = input.Password

	if err := u.BeforeAddUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process user data"})
		return
	}

	_, err := Models.AddUser(db, &u)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"registration success"}) 

}