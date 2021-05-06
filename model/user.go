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
	// get collection
	databaseName := config.DBcfg[pgpFileConnName].Database
	userCollection := tool.GetClient(pgpFileConnName).Database(databaseName).Collection(userCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// simple convertion from []PGPFile to []interface{}
	document := make([]interface{}, len(users))
	for i, user := range users {
		document[i] = user
	}

	// insert documents (converted from input)
	res, err := userCollection.InsertMany(ctx, document)
	if err != nil {
		errmsg := "arapgp.model.user => InsertUsers: collection InsertMany failed;"
		log.WithFields(log.Fields{"res": res, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return nil
}

// GetUsers is a "R"ead Operation
func GetUsers(users []User, filter bson.D) (err error) {
	// get collection
	databaseName := config.DBcfg[userConnName].Database
	userCollection := tool.GetClient(userConnName).Database(databaseName).Collection(userCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// generate cursor
	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		errmsg := "arapgp.model.user => GetUser: collection Find failed"
		log.WithFields(log.Fields{"cur": cur, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	defer cur.Close(ctx)

	// cursor get all
	if err = cur.All(ctx, &users); err != nil {
		errmsg := "arapgp.model.user => GetUser: cursor.All failed"
		log.WithFields(log.Fields{"cur": cur, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return nil
}

// UpdateUsers will update all Users that get through filter
func UpdateUsers(update bson.M, filter bson.D) (err error) {
	// get collection
	databaseName := config.DBcfg[pgpFileConnName].Database
	userCollection := tool.GetClient(pgpFileConnName).Database(databaseName).Collection(userCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// update documents
	res, err := userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		errmsg := "arapgp.model.user => UpdateUsers: collection UpdateMany failed;"
		log.WithFields(log.Fields{"res": res, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return nil
}

// DeleteUsers will delete all Users that get through filter
func DeleteUsers(filter bson.D) (err error) {
	// get collection
	databaseName := config.DBcfg[pgpFileConnName].Database
	userCollection := tool.GetClient(pgpFileConnName).Database(databaseName).Collection(userCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// delete documents
	res, err := userCollection.DeleteMany(ctx, filter)
	if err != nil {
		errmsg := "arapgp.model.user => DeleteUsers: collection DeleteMany failed;"
		log.WithFields(log.Fields{"res": res, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return nil
}
