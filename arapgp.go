package main

import (
	"strconv"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/route"
)

func main() {
	config.Setup("./arapgp.server.json")
	router := route.InitRouter()

	router.Use(route.CORS())
	router.Run(":" + strconv.Itoa(config.Svccfg.Port))
}
