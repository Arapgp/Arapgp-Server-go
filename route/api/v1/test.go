package v1

import (
	"net/http"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/tool"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	pgpFileConnName       = "mongo"
	pgpFileCollectionName = "pgpfile"
	userConnName          = "mongo"
	userCollectionName    = "user"
)

// Ping is to check Back-end & Database
func Ping(c *gin.Context) {
	pgpFileCollection := tool.GetClient(pgpFileConnName).Database(config.DBcfg[pgpFileConnName].Database).Collection(pgpFileCollectionName)
	userCollection := tool.GetClient(pgpFileConnName).Database(config.DBcfg[userConnName].Database).Collection(userCollectionName)

	log.WithFields(log.Fields{
		"pgpfile": pgpFileCollection, "user": userCollection,
	}).Infoln("All Collections are Healthy!")

	c.JSON(http.StatusOK, gin.H{"status": "OK", "ping": "pong"})
	return
}
