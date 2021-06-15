package v1

import (
	"net/http"

	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetPubKey is to "Get PubKey" by User's name
func GetPubKey(c *gin.Context) {
	queryName := c.Query("username")

	users := make([]model.User, 1)
	err := model.GetUsers(users, bson.M{"profile.name": queryName})
	if err != nil || users[0].Profile.Name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "PubKey does not exist!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "pubKey": users[0].PubKey})
	return
}

// PostPutPubKey is to "Post PubKey" by User's uid
// POST & PUT is the same
// need session
func PostPutPubKey(c *gin.Context) {
	var json JSONPubKey
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad post request!"})
		return
	}

	// needn't process error, because auth has done these job
	session, _ := c.Cookie("SeesionId")

	// update PubKey
	err := model.UpdateUsers(
		bson.M{"$set": bson.M{
			"pubkey": json.PubKey,
		}},
		bson.D{{Key: "session", Value: session}},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}

// DeletePubKey is to "Delete PubKey" by User's uid
// need session
func DeletePubKey(c *gin.Context) {
	// needn't process error, because auth has done these job
	session, _ := c.Cookie("SeesionId")

	// update PubKey
	err := model.UpdateUsers(
		bson.M{"$set": bson.M{
			"pubkey": "",
		}},
		bson.D{{Key: "session", Value: session}},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "Unexpected error!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
	return
}

// JSONPubKey is used in "/api/v1/pubKey POST, PUT"
type JSONPubKey struct {
	PubKey string `json:"pubKey" binding:"required"`
}
