package test

import (
	"strconv"
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/route"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var urlBase string

func setupConfig(t *testing.T) {
	config.Setup("../arapgp.server.json")
	urlBase = "http://" + config.Svccfg.Host + ":" + strconv.Itoa(config.Svccfg.Port)

	log.Info("Setup, Test Begin, url base is:", urlBase)
}

func teardownConfig(t *testing.T) {
	config.Teardown()
	urlBase = ""
	log.Info("Teardown, Test End")
}

func setupRouter(t *testing.T) (r *gin.Engine) {
	return route.InitRouter()
}
