package main

import (
	"strconv"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/route"
	"github.com/gin-contrib/cors"
)

func main() {
	config.Setup("./arapgp.server.json")
	router := route.InitRouter()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"Access-Control-Allow-Origin"},
		MaxAge:           12 * time.Hour,
	}))

	router.Run(":" + strconv.Itoa(config.Svccfg.Port))
}
