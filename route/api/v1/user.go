package v1

import (
	"log"
	"net/http"

	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Register is to register a new user
func Register(c *gin.Context) {
	var json JSONUsernamePassword
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad JSON Post!" + err.Error()})
		return
	}
	log.Println(json)

	// len(Username) must be less than 50
	if len(json.Username) > 50 || len(json.Username) <= 0 {
		c.JSON(http.StatusAccepted, gin.H{"error": "Username not legal!"})
		return
	}

	// check whether user exists
	users := []model.User{}
	err := model.GetUsers(users, bson.D{{Key: "name", Value: json.Username}})
	if len(users) != 0 || err != nil {
		c.JSON(http.StatusAccepted, gin.H{"error": "Username already exists!"})
		return
	}

	// do register(insert) job
	users = []model.User{{
		Profile: model.UserProfile{Name: json.Username, Password: json.Password},
		Files:   nil,
		PubKey:  "",
	}}
	err = model.InsertUsers(users)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"error": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "OK"})
	return
}

// Login is a function that process login
// 1. check username & password
// 2. return result
func Login(c *gin.Context) {
	return
}

// Logout is to logout(
func Logout(c *gin.Context) {
	return
}

// GetUsersByName is to "Get UserList" by Name-prefix
func GetUsersByName(c *gin.Context) {
	return
}

// JSONUsernamePassword is a type for login/register api
type JSONUsernamePassword struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
