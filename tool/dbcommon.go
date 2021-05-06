package tool

import (
	"context"
	"strconv"

	"github.com/Arapgp/Arapgp-Server-go/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetClient return a MongoDB Client
func GetClient(ConnName string) *mongo.Client {
	// compose mongodb uri
	connCfg := config.DBcfg[ConnName]
	uri := "mongodb://" + connCfg.Host + ":" + strconv.Itoa(connCfg.Port)

	// init client options & get client
	clientOpt := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOpt)
	if err != nil {
		log.WithFields(log.Fields{
			"uri": uri, "opt": clientOpt, "client": client,
		}).Fatalln("tool.GetClient NewClient failed")
	}

	// try to connect mongodb using client
	err = client.Connect(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"opt": clientOpt, "client": client,
		}).Fatalln("tool.GetClient Connect failed")
	}
	return client
}
