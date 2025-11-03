package test

import (
	"testing"

	"github.com/silentryan/network-test/files"
)

func TestMigrateData(t *testing.T) {
	// X is empty, as the
	var cols []string = []string{"A", "B", "D", "C", "E", "X", "X", "S", "G"}
	originData := files.GetContentByCols("C:/Users/L115840/Desktop/FeitiankeyCount.xlsx", "Sheet1", cols)

	files.WriteDatetoTable("C:/Users/L115840/Desktop/feitian-key.xlsx", "key", originData)
}
