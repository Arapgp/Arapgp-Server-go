package test

import (
	"testing"

	"github.com/Arapgp/Arapgp-Server-go/pkg/sfs"
	"github.com/stretchr/testify/assert"
)

func Test_Sfs(t *testing.T) {
	ast := assert.New(t)

	err := sfs.WriteContentByPath("./", "1.txt", "hahahah")
	ast.Empty(err)

	err = sfs.WriteContentByPath("./", "2.txt", "hohohoh")
	ast.Empty(err)

	content, err := sfs.GetContentByPath("./", "2.txt")
	ast.Empty(err)
	ast.NotEmpty(content)

	err = sfs.DeleteFileByPath("./", "1.txt")
	ast.Empty(err)

	content, err = sfs.GetContentByPath("./", "1.txt")
	ast.NotEmpty(err)
	ast.Empty(content)

	err = sfs.WriteContentByPath("./sfs_test/", "3.txt", "heheheh")
	ast.Empty(err)
}
