package v1

import (
	"net/http"

	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// PostFileByUserName is to "POST/PUT PGPFile" by username & file name
func PostFileByUserName(c *gin.Context) {

}

// PutFileByUserName is to "PUT PGPFile" by username & file name
func PutFileByUserName(c *gin.Context) {

}

// GetFileByUserName is to "GET PGPFile" by username & file name
func GetFileByUserName(c *gin.Context) {

}

// DeleteFileByUserName is to "DELETE PGPFile" by username & file name
func DeleteFileByUserName(c *gin.Context) {

}

// GetFilesInfoByUserName is to "GET PGPFiles' info" by username
func GetFilesInfoByUserName(c *gin.Context) {
	username := c.Param("username")

	users := make([]model.User, 1)
	err := model.GetUsers(users, bson.D{{Key: "profile.name", Value: username}})
	if users[0].Profile.Name == "" || err != nil {
		log.WithFields(
			log.Fields{"username": username, "user": users, "err": err.Error()},
		).Warningln("api.v1.GetFilesInfoByUserName err: user do not exist")
		c.JSON(http.StatusOK, gin.H{"status": "User do not exist!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"info":   users[0].Files,
	})
	return
}

// JSONPostPutFile is a json for Post / Put File
type JSONPostPutFile struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}
