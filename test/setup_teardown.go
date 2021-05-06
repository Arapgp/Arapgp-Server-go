package test

import (
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/config"
	log "github.com/sirupsen/logrus"
)

func setupConfig(t *testing.T) {
	config.Setup("../arapgp.server.json")
	log.Info("Setup, Test Begin")
}

func teardownConfig(t *testing.T) {
	config.Teardown()
	log.Info("Teardown, Test End")
}
