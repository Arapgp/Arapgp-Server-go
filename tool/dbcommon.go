package tool

import (
	"context"
	"fmt"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GetClient return a MongoDB Client.
// This function gets client via uri defined in arapgp.server.json,
// then, tries to connect target database,
// and ping it to make sure connection built successfully.
func GetClient(ConnName string) *mongo.Client {
	// compose mongodb uri
	connCfg := config.DBcfg[ConnName]
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		connCfg.Username, connCfg.Password,
		connCfg.Host, connCfg.Port, connCfg.Database,
	)

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Fatalln("tool.GetClient Ping failed")
	}
	return client
}
