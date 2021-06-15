package v1

import (
	"net/http"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/Arapgp/Arapgp-Server-go/pkg/sfs"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// PostFileByUserName is to "POST PGPFile" by username & file name
func PostFileByUserName(c *gin.Context) {
	username := c.Param("username")
	var json JSONPostPutFile
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad" + c.Request.Method + "request!"})
		return
	}

	// get users PubKey(insert pgpfile needed)
	// check PubKey: (1) is null? (2) is matched?
	users := make([]model.User, 1)
	err := model.GetUsers(users, bson.M{"profile.name": username})
	if err != nil || users[0].Profile.Name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "User does not exist!"})
		return
	}
	if users[0].PubKey != json.PubKey || users[0].PubKey == "" {
		c.JSON(http.StatusOK, gin.H{"status": "Public Key do not qualified!"})
		return
	}

	// check whether file already exists
	checkFiles := make([]model.PGPFile, 1)
	err = model.GetPGPFiles(checkFiles, bson.M{"name": json.Name, "author": username})
	if err != nil || checkFiles[0].Name != "" {
		c.JSON(http.StatusOK, gin.H{"status": "File already exists!"})
		return
	}

	// create file in archive folder
	err = sfs.WriteContentByPath(config.Svccfg.File+username+"/", json.Name, json.Content)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	// a new file
	files := []model.PGPFile{
		{
			Name: json.Name, Author: username, Size: len(json.Content),
			CreateTime: time.Now(), LastModifyTime: time.Now(),
			Path: "/" + username + "/" + json.Name, PubKey: json.PubKey,
		},
	}
	err = model.InsertPGPFiles(files)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	// insert documents to user.files here
	// model.InsertPGPFiles aims to insert documents to pgpfile collection
	err = model.UpdateUsers(
		bson.M{"$push": bson.M{
			"files": bson.M{"$each": files},
		}},
		bson.M{"profile.name": username},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}

// PutFileByUserName is to "PUT PGPFile" by username & file name
func PutFileByUserName(c *gin.Context) {
	username := c.Param("username")
	var json JSONPostPutFile
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad" + c.Request.Method + "request!"})
		return
	}

	// get users PubKey(insert pgpfile needed)
	// check PubKey: (1) is null? (2) is matched?
	users := make([]model.User, 1)
	err := model.GetUsers(users, bson.M{"profile.name": username})
	if err != nil || users[0].Profile.Name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "User does not exist!"})
		return
	}
	if users[0].PubKey != json.PubKey || users[0].PubKey == "" {
		c.JSON(http.StatusOK, gin.H{"status": "Public Key do not qualified!"})
		return
	}

	// check whether file already exists
	checkFiles := make([]model.PGPFile, 1)
	err = model.GetPGPFiles(checkFiles, bson.M{"name": json.Name, "author": username})
	if err != nil || checkFiles[0].Name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "File do not exist!"})
		return
	}

	// update file in archive folder
	err = sfs.WriteContentByPath(config.Svccfg.File+username+"/", json.Name, json.Content)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	// update document in pgpfile collection
	err = model.UpdatePGPFiles(
		bson.M{"$set": bson.M{
			"lastmodifytime": time.Now(),
		}},
		bson.M{"author": username, "name": json.Name},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	// insert documents to user.files here
	// model.InsertPGPFiles aims to insert documents to pgpfile collection
	err = model.UpdateUsers(
		bson.M{"$set": bson.M{
			"files.$.lastmodifytime": time.Now(),
		}},
		bson.M{"profile.name": username},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
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
	PubKey  string `json:"pubKey" binding:"required"`
}

// JSONGetFile is a json for Get File
type JSONGetFile struct {
	Name string `json:"name" binding:"required"`
}
