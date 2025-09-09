package Controllers

import (
	"net/http"

	"database/sql"
  	
	"github.com/gin-gonic/gin"
	
	"example/Wave_Seekers_Back/Models"

	"example/Wave_Seekers_Back/Utils/Token"
)

var db *sql.DB

func InitializeDB(database *sql.DB) {
	db = database
}

func CurrentUser(c *gin.Context){

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	u, err := Models.GetCurrentUserByID(db, user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","data":u})
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

	userID, err := Models.AddUser(db, &u)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := token.GenerateToken(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful", 
		"token": token,
		"user_id": userID,
	})
}

func Logout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Successfully logged out. Please remove token on client side.",
    })
}