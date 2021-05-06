package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

var (
	// Svccfg is an entry to ServerConfigModel
	Svccfg *ServerConfigModel = &ServerConfigModel{}
	// DBcfg maintains several Database config info
	DBcfg DatabaseConfigModels = make(DatabaseConfigModels, 0)
)

// DatabaseConfigModel (DBcfg), model of "db" part in arapgp.server.ini
type DatabaseConfigModel struct {
	Host     string
	Port     int
	Username string
	Password string
}

// DatabaseConfigModels is map(string => DatabaseConfigModel)
type DatabaseConfigModels []DatabaseConfigModel

// ServerConfigModel (Svccfg), model of "server" part in arapgp.server.ini
type ServerConfigModel struct {
	Name string
	Port int
}

// Setup is to init config models
/*
func Setup() {
	res, err := cfg.ReadConfigFile("arapgp.server.json")
	cfg.MapConfig(res)
}
*/

// Unmarshal is to implement interface "Unmarshaler"
func (cfg *DatabaseConfigModel) Unmarshal(res gjson.Result) {
	cfg.Host = res.Get("host").String()
	cfg.Port = int(res.Get("port").Int())
	cfg.Username = res.Get("username").String()
	cfg.Password = res.Get("password").String()
}

// Unmarshal is to implement interface "Unmarshaler"
func (cfg *ServerConfigModel) Unmarshal(res gjson.Result) {
	cfg.Name = res.Get("name").String()
	cfg.Port = int(res.Get("port").Int())
}

// Unmarshal is to implement interface "Unmarshaler"
func (cfg *DatabaseConfigModels) Unmarshal(res gjson.Result) {
	// DatabaseConfigModels is a slice
	for arr, i := res.Array(), 0; i < len(arr); i++ {
		// cfg is a slice, need append(cfg, DatabaseConfigModel)
		tmp := &DatabaseConfigModel{}
		tmp.Unmarshal(arr[i])
		log.WithFields(log.Fields{"tmp": tmp, "arr[i]": arr[i]}).Info("after unmarshal arr[i] to tmp")
		*cfg = append(*cfg, *tmp)
	}
	log.WithFields(log.Fields{"cfg": cfg}).Info("after unmarshal of databaseconfigmodels")
}
