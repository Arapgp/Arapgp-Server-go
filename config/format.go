package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// variables about config
var (
	// DBcfg maintains several Database config info
	DBcfg DatabaseConfigModels = make(DatabaseConfigModels)
	// Svccfg is an entry to ServerConfigModel
	Svccfg *ServerConfigModel = &ServerConfigModel{}
)

// DatabaseConfigModel (DBcfg), model of "db" part in arapgp.server.ini
type DatabaseConfigModel struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

// DatabaseConfigModels is map(string => DatabaseConfigModel)
type DatabaseConfigModels map[string]DatabaseConfigModel

// ServerConfigModel (Svccfg), model of "server" part in arapgp.server.ini
type ServerConfigModel struct {
	Name string
	Port int
}

// Unmarshal is to implement interface "Unmarshaler"
func (cfg *DatabaseConfigModel) Unmarshal(res gjson.Result) {
	cfg.Host = res.Get("host").String()
	cfg.Port = int(res.Get("port").Int())
	cfg.Username = res.Get("username").String()
	cfg.Password = res.Get("password").String()
	cfg.Database = res.Get("database").String()
}

// Unmarshal is to implement interface "Unmarshaler"
func (cfg *ServerConfigModel) Unmarshal(res gjson.Result) {
	cfg.Name = res.Get("name").String()
	cfg.Port = int(res.Get("port").Int())
}

// Unmarshal is to implement interface "Unmarshaler"
func (cfg DatabaseConfigModels) Unmarshal(res gjson.Result) {
	// DatabaseConfigModels is a map
	for k, v := range res.Map() {
		tmp := DatabaseConfigModel{}
		tmp.Unmarshal(v)
		log.WithFields(log.Fields{"tmp": tmp, "v": v}).Info("DatabaseConfigModels Unmarshal: after unmarshal v to tmp")
		cfg[k] = tmp
	}
	log.WithFields(log.Fields{"cfg": cfg}).Info("DatabaseConfigModels Unmarshal: end")
}
