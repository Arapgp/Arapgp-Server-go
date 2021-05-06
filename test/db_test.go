package test

import (
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/config"
	"github.com/Arapgp/Arapgp-Server-go/tool"
)

func Test_ConnectDatabase(t *testing.T) {
	config.Setup("../arapgp.server.json")
	tool.GetClient("mongo")
}
