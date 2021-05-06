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
	Name   string
	Author string
	Size   int

	// Create_time represents when file sent to server
	CreateTime     time.Time
	LastModifyTime time.Time

	// Path means the place where this "File" stored
	// and this path is a relative path.
	// absolute path = path prefix + path (relative path)
	Path string
}

// InsertPGPFiles is to insert multi-files
func InsertPGPFiles(files []PGPFile) (err error) {
	databaseName := config.DBcfg[pgpFileConnName].Database
	pgpFileCollection := tool.GetClient(pgpFileConnName).Database(databaseName).Collection(pgpFileCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// simple convertion from []PGPFile to []interface{}
	document := make([]interface{}, len(files))
	for i, file := range files {
		document[i] = file
	}

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
	return
}

// UpdatePGPFiles will update all PGPFiles that get through filter
func UpdatePGPFiles(update bson.M, filter bson.D) (err error) {
	return
}

// DeletePGPFiles will delete all PGPFiles that get through filter
func DeletePGPFiles(filter bson.D) (err error) {
	return
}
