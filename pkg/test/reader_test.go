package test

import (
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/pkg/cfg"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_reader(t *testing.T) {
	ast := assert.New(t)

	// get gjson.Result from config file
	res, err := cfg.ReadConfigFile("arapgp.server.json")
	ast.NotEmpty(res)
	log.WithFields(log.Fields{"Res": res}).Info("After Read Config File")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// m as MapConfig function parameter
	var m map[string]cfg.Unmarshaler = map[string]cfg.Unmarshaler{
		"db":     config.DBcfg,
		"server": config.Svccfg,
	}

	err = cfg.MapConfig(res, m)
	ast.NotEmpty(m["db"])
	ast.NotEmpty(m["server"])
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	db := config.DBcfg
	svc := config.Svccfg
	ast.Equal(svc, &config.ServerConfigModel{Name: "arapgp", Port: 3000})
	ast.Equal(db, config.DatabaseConfigModels{"mongo": {Host: "127.0.0.1", Port: 27017, Username: "ljg", Password: "ljg", Database: "ljgtest"}})
}
