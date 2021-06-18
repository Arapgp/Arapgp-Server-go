package main

import (
	"strconv"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/route"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	config.Setup("./arapgp.server.json")
	router := route.InitRouter()

	// CORS
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, POST, PUT, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		MaxAge:         12 * time.Hour,
	}))

	router.Run(":" + strconv.Itoa(config.Svccfg.Port))
}
