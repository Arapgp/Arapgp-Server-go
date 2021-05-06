package test

import (
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_GetUser(t *testing.T) {
	ast := assert.New(t)
	setupConfig(t)
	defer teardownConfig(t)

	users := []model.User{}
	err := model.GetUsers(users, bson.D{})
	ast.Empty(err)

	log.WithFields(log.Fields{"users": users}).Info("GetUser end")
}
