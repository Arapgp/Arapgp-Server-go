package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/Arapgp/Arapgp-Server-go/pkg/session"
	"github.com/Arapgp/Arapgp-Server-go/pkg/shatool"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Register is to register a new user
func Register(c *gin.Context) {
	var json JSONUsernamePassword
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad json post!"})
		return
	}

	// len(Username) must be less than 50
	if len(json.Username) > 50 || len(json.Username) <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": "Username not legal!"})
		return
	}

	// check whether user exists
	users := make([]model.User, 1)
	err := model.GetUsers(users, bson.M{"profile.name": json.Username})
	if users[0].Profile.Name == json.Username || err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Username already exists!"})
		return
	}

	// do register(insert) job
	password := shatool.Sha256String(json.Password)
	users = []model.User{{
		Profile: model.UserProfile{Name: json.Username, Password: password, LastLoginTime: time.Now()},
		Files:   []model.PGPFile{},
		PubKey:  "",
	}}
	err = model.InsertUsers(users)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}

// Login is a function that process login
// 1. check username & password
// 2. update last login time of user
// 3. return result{ status lastLoginTime session }
func Login(c *gin.Context) {
	// bind request json
	var json JSONUsernamePassword
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad json post!"})
		return
	}

	// get only 1 user from database via model.GetUsers
	users := make([]model.User, 1)
	err := model.GetUsers(users, bson.D{{Key: "profile.name", Value: json.Username}})
	// can only use `users[0].Profile.Name == ""` to check whether GetUsers failed
	// how terrible golang and mongo-go-driver are!
	if err != nil || users[0].Profile.Name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "User not existed!"})
		return
	}
	lastLoginTime := users[0].Profile.LastLoginTime

	// check username & password
	if pwd := shatool.Sha256String(json.Password); users[0].Profile.Password != pwd {
		log.Println(pwd, users[0].Profile.Password)
		c.JSON(http.StatusOK, gin.H{"status": "Username or password wrong!"})
		return
	}

	// update last login time
	newSession := session.GenerateSession(json.Username)
	err = model.UpdateUsers(
		bson.M{"$set": bson.M{
			"profile.lastlogintime": time.Now(),
			"session":               newSession,
		}},
		bson.M{"profile.name": json.Username},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "User not existed!"})
		return
	}

	// OK, return
	c.JSON(http.StatusOK, gin.H{
		"status": "OK", "session": newSession,
		"lastLoginTime": lastLoginTime,
	})
	return
}

// Logout is to logout
// need session
func Logout(c *gin.Context) {
	// get session, Auth already, no error
	session, _ := c.Cookie("SessionId")
	err := model.UpdateUsers(
		bson.M{"$set": bson.M{"session": ""}},
		bson.M{"session": session},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}

// GetUsersByName is to "Get UserList" by Name-prefix
// supposed to provide session, but now needn't
func GetUsersByName(c *gin.Context) {
	// get params
	query, ok := c.GetQuery("query")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	// only return 10 users,
	// and in mongodb-go-driver, "key": /^val/ is forbidden.
	// because it needn't `""` in mongodb shell, Find will view "/" as normal slash.
	// use $regex: "^val" instead.
	users := make([]model.User, 10)
	err := model.GetUsers(users, bson.M{"profile.name": bson.M{"$regex": "^" + query}})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	// get userList
	userList := JSONUserList{}
	for _, user := range users {
		if user.Profile.Name == "" {
			break
		}
		userList = append(userList, JSONUser{Username: user.Profile.Name})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"userList": userList,
	})
	return
}

// JSONUsernamePassword is a type for login/register api
type JSONUsernamePassword struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// JSONStatus is to "Get"
type JSONStatus struct {
	Status string `json:"status" binding:"required"`
}

// JSONUser is used when GET /api/v1/user
type JSONUser struct {
	Username string
}

// JSONUserList is used when GET /api/v1/user
type JSONUserList []JSONUser

// JSONGetUser is used as response when GET /api/v1/user
type JSONGetUser struct {
	Status   string       `json:"status" binding:"required"`
	UserList JSONUserList `json:"userList" binding:"required"`
}
