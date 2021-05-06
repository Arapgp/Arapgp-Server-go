package config

import (
	"github.com/Arapgp/Arapgp-Server-go/pkg/cfg"
	log "github.com/sirupsen/logrus"
)

// Setup is a tool function to setup environment.
// 1. init app via arapgp.server.json, load DBcfg & Svccfg
// This exported function need called at the beginning
func Setup(svcfgPath string) {
	setupDBSvc(svcfgPath)
}

func setupDBSvc(path string) {
	// Get gjson.Result from config file
	res, err := cfg.ReadConfigFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	// models as MapConfig function parameter
	models := map[string]cfg.Unmarshaler{
		"db":     DBcfg,
		"server": Svccfg,
	}

	// Map res(gjson.Result) To DBcfg & Svccfg
	err = cfg.MapConfig(res, models)
	if err != nil {
		log.Fatal(err.Error())
	}
}
