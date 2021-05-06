package model

import (
	"context"
	"errors"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/tool"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	userConnName       = "mongo"
	userCollectionName = "user"
)

// User as a document
type User struct {
	Profile UserProfile
	Files   []PGPFile
	PubKey  string
}

// UserProfile is a type used in User
type UserProfile struct {
	// Name is one of properties to identify users
	Name string
	// Password need stored after processing
	Password string
}

// InsertUsers is to insert multi-users
func InsertUsers(users []User) (err error) {
	return
}

// GetUsers is a "R"ead Operation
func GetUsers(users []User, filter bson.D) (err error) {
	databaseName := config.DBcfg[userConnName].Database
	userCollection := tool.GetClient(userConnName).Database(databaseName).Collection(userCollectionName)
	log.WithFields(log.Fields{"dbName": databaseName, "colName": userCollectionName, "connName": userConnName}).Info("???")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		errmsg := "arapgp.model.user => GetUser: collection Find failed"
		log.WithFields(log.Fields{"cur": cur, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &users); err != nil {
		errmsg := "arapgp.model.user => GetUser: cursor.All failed"
		log.WithFields(log.Fields{"cur": cur, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	log.Infoln(users)
	return nil
}

// UpdateUsers will update all Users that get through filter
func UpdateUsers(update bson.M, filter bson.D) (err error) {
	return
}

// DeleteUsers will delete all Users that get through filter
func DeleteUsers(filter bson.D) (err error) {
	return
}
