package cfg

import (
	"errors"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// MapConfig is a tool function wrap MapTo in go-ini
// read config in .ini, map them to given struct
func MapConfig(cfg gjson.Result, models map[string]Unmarshaler) (err error) {
	for k, v := range models {
		res := cfg.Get(k)
		v.Unmarshal(res)
		if err != nil {
			log.WithFields(log.Fields{"key": k, "val": v}).Warningln(err.Error())
			return errors.New("arapgp.pkg.config.reader => MapConfig error")
		}
	}
	return nil
}

// ReadConfigFile is a tool function
// "path" is path to *.json file
// get gjson.Result through *.json path
func ReadConfigFile(path string) (res gjson.Result, err error) {
	fd, err := os.Open(path)
	if err != nil {
		dir, _ := os.Getwd()
		log.WithFields(log.Fields{
			"pwd": dir,
			"err": err.Error(),
		}).Fatalln("arapgp.pkg.cfg reader.go => ReadConfigFile error: os.Open failed")
		return gjson.Result{}, err
	}

	contents, err := ioutil.ReadAll(fd)
	if err != nil {
		log.WithFields(log.Fields{
			"contents": contents,
			"err":      err.Error(),
		}).Fatalln("arapgp.pkg.cfg reader.go => ReadConfigFile error: ioutil.ReadAll failed")
		return gjson.Result{}, err
	}

	return gjson.Parse(string(contents)), nil
}
