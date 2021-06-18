package main

import (
	"strconv"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/route"
	"github.com/gin-contrib/cors"
)

func main() {
	config.Setup("./arapgp.server.json")
	router := route.InitRouter()

	// CORS
	router.Use(cors.Default())

	router.Run(":" + strconv.Itoa(config.Svccfg.Port))
}
