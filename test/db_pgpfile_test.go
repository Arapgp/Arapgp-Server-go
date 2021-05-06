package test

import (
	"testing"
	"time"

	"github.com/Arapgp/Arapgp-Server-go/model"
	"github.com/stretchr/testify/assert"
)

func Test_InsertPGPFiles(t *testing.T) {
	ast := assert.New(t)
	setupConfig(t)
	defer teardownConfig(t)

	tcases := []struct {
		name  string
		files []model.PGPFile
	}{
		{
			name: "InsertPGPFiles_1",
			files: []model.PGPFile{
				{Name: "ljg", Author: "skyleaworlder", Size: 114514, CreateTime: time.Now(), LastModifyTime: time.Now(), Path: "/ljg"},
				{Name: "gjl", Author: "skyleaworlder", Size: 415411, CreateTime: time.Now(), LastModifyTime: time.Now(), Path: "/ljg"},
			},
		},
		{
			name: "InsertPGPFiles_2",
			files: []model.PGPFile{
				{Name: "htc", Author: "eol", Size: 114514, CreateTime: time.Now().Add(time.Minute), LastModifyTime: time.Now().Add(time.Hour), Path: "/eol"},
			},
		},
	}

	for _, c := range tcases {
		t.Run(c.name, func(t *testing.T) {
			err := model.InsertPGPFiles(c.files)
			ast.Empty(err)
		})
	}
}
