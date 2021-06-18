package route

import (
	"net/http"

	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	userConnName       = "mongo"
	userDatabaseName   = "ljgtest"
	userCollectionName = "user"
)

// Auth is a wrapper to check authentication
func Auth(inner gin.HandlerFunc) (outer gin.HandlerFunc) {
	return func(c *gin.Context) {
		// check cookie in request
		session, err := c.Cookie("SessionId")
		if err != nil {
			log.WithFields(log.Fields{
				"session": session,
				"err":     err.Error(),
			}).Warningln("arapgp.route.wrapper => Auth *gin.Context.Cookie failed.")
			c.JSON(http.StatusOK, gin.H{"status": "Unauthenticated!"})
			return
		}

		// check session in database
		users := make([]model.User, 1)
		err = model.GetUsers(users, bson.M{"session": session})
		// use users[0].Profile.Name != "" to verify whether found
		if err != nil || users[0].Profile.Name == "" {
			log.WithFields(log.Fields{
				"users": users,
				"err":   err.Error(),
			}).Warningln("arapgp.route.wrapper => Auth model.GetUser failed.")
			c.JSON(http.StatusOK, gin.H{"status": "Unauthenticated!"})
			return
		}

		// execute inner function
		inner(c)
		return
	}
}

// HeaderWrapper is a wrapper to add header
func HeaderWrapper(inner gin.HandlerFunc) (outer gin.HandlerFunc) {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		// execute inner function
		inner(c)
		return
	}
}
