package main

import (
	"github.com/Arapgp/Arapgp-Server-go/route"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := route.InitRouter()
	log.Info("route:", router)
}
