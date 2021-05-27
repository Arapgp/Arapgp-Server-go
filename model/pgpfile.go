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
	pgpFileConnName       = "mongo"
	pgpFileCollectionName = "pgpfile"
)

// PGPFile is file document type
type PGPFile struct {
	// Name is used to figure out src & dst.
	// namely, "repeat name" is forbidden
	Name   string `bson:"name"`
	Author string `bson:"author"`
	Size   int    `bson:"size"`

	// Create_time represents when file sent to server
	CreateTime     time.Time `bson:"createtime"`
	LastModifyTime time.Time `bson:"lastmodifytime"`

	// Path means the place where this "File" stored
	// and this path is a relative path.
	// absolute path = path prefix + path (relative path)
	Path   string `bson:"path"`
	PubKey string `bson:"pubkey"`
}

// InsertPGPFiles is to insert multi-files
func InsertPGPFiles(files []PGPFile) (err error) {
	// get collection
	databaseName := config.DBcfg[pgpFileConnName].Database
	pgpFileCollection := tool.GetClient(pgpFileConnName).Database(databaseName).Collection(pgpFileCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// simple convertion from []PGPFile to []interface{}
	document := make([]interface{}, len(files))
	for i, file := range files {
		document[i] = file
	}

	// insert documents (converted from input)
	res, err := pgpFileCollection.InsertMany(ctx, document)
	if err != nil {
		errmsg := "arapgp.model.pgpfile => InsertPGPFiles: collection InsertMany failed;"
		log.WithFields(log.Fields{"res": res, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return nil
}

// GetPGPFiles will get Many PGPfiles from mongo.Collection
func GetPGPFiles(files []PGPFile, filter bson.D) (err error) {
	// get collection
	databaseName := config.DBcfg[userConnName].Database
	pgpFileCollection := tool.GetClient(userConnName).Database(databaseName).Collection(pgpFileCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// generate cursor
	cur, err := pgpFileCollection.Find(ctx, filter)
	if err != nil {
		errmsg := "arapgp.model.pgpfile => GetPGPFiles: collection Find failed"
		log.WithFields(log.Fields{"cur": cur, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	defer cur.Close(ctx)

	// cursor get all
	if err = cur.All(ctx, &files); err != nil {
		errmsg := "arapgp.model.pgpfile => GetPGPFile: cursor.All failed"
		log.WithFields(log.Fields{"cur": cur, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return
}

// UpdatePGPFiles will update all PGPFiles that get through filter
func UpdatePGPFiles(update bson.M, filter bson.D) (err error) {
	// get collection
	databaseName := config.DBcfg[userConnName].Database
	pgpFileCollection := tool.GetClient(userConnName).Database(databaseName).Collection(pgpFileCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// update documents
	res, err := pgpFileCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		errmsg := "arapgp.model.pgpfile => UpdatePGPFiles: collection UpdateMany failed;"
		log.WithFields(log.Fields{"res": res, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return
}

// DeletePGPFiles will delete all PGPFiles that get through filter
func DeletePGPFiles(filter bson.D) (err error) {
	// get collection
	databaseName := config.DBcfg[userConnName].Database
	pgpFileCollection := tool.GetClient(userConnName).Database(databaseName).Collection(pgpFileCollectionName)

	// generate context
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// delete documents
	res, err := pgpFileCollection.DeleteMany(ctx, filter)
	if err != nil {
		errmsg := "arapgp.model.pgpfile => DeletePGPFiles: collection DeleteMany failed;"
		log.WithFields(log.Fields{"res": res, "err": err.Error()}).Warningln(errmsg)
		return errors.New(errmsg + err.Error())
	}
	return
}
