package test

import (
	"context"
	"testing"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/tool"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_ConnectDatabase(t *testing.T) {
	ast := assert.New(t)
	setupConfig(t)
	defer teardownConfig(t)

	tcases := []struct {
		name     string
		connName string
		dbName   string
		colName  string
	}{
		{name: "ConnectDatabase_1", connName: "mongo", dbName: "ljgtest", colName: "pgpfiles"},
		{name: "ConnectDatabase_2", connName: "mongo", dbName: "ljgtest", colName: "user"},
	}

	for _, c := range tcases {
		t.Run(c.name, func(t *testing.T) {
			client := tool.GetClient(c.connName)
			ast.NotEmpty(client)
			collection := client.Database(c.dbName).Collection(c.colName)
			ast.NotEmpty(collection)

			log.WithFields(log.Fields{
				"collection": collection,
			}).Info("get Database and Collection")
		})
	}
}

func Test_SelectDatabase(t *testing.T) {
	ast := assert.New(t)
	setupConfig(t)
	defer teardownConfig(t)

	tcases := []struct {
		name     string
		connName string
		dbName   string
		colName  string
	}{
		{name: "SelectDatabase_1", connName: "mongo", dbName: "ljgtest", colName: "pgpfiles"},
		{name: "SelectDatabase_2", connName: "mongo", dbName: "ljgtest", colName: "user"},
	}

	for _, c := range tcases {
		t.Run(c.name, func(t *testing.T) {
			// get collection
			client := tool.GetClient(c.connName)
			ast.NotEmpty(client)
			collection := client.Database(c.dbName).Collection(c.colName)
			ast.NotEmpty(collection)

			// generate ctx, and then get cursor
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			cur, err := collection.Find(ctx, bson.D{})
			ast.Empty(err)
			defer cur.Close(ctx)

			// use cursor to scan collection
			for cur.Next(ctx) {
				var result bson.D
				err := cur.Decode(&result)
				ast.Empty(err)
				log.WithFields(log.Fields{"result": result}).Info("cur.Decode")
			}
			ast.Empty(cur.Err())
		})
	}
}
