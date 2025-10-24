package test

import (
	"testing"

	"github.com/silentryan/network-test/output"
)

func TestCrateFile(t *testing.T) {
	// running success
	// file must end with `.xlsx`, or error: unsupported workbook file format
	output.CreatExcel("C:/Users/L115840/Desktop/workspace/network-test/test/testexcel.xlsx")
}
