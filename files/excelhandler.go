package files

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// handle a input a excel

// This struct should be group of what informantion we want to get from a excel
// Name: Excel file path, return a workbook
// Sheet: Excel may should multiple sheet, this sheet is the one we want
// Table: A unique table of where we get information from
// Headers: Map[Cols]Target
type ExlHandler struct {
	sheet   string
	table   string
	headers []string
}

func Excelhandle(filename string) error {
	excel, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}

	return nil
}

// Read All tables in a Excel
func GetAllTables() []string {

	return nil
}
