package output

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// According to object create a excel table
// sheet1 will default created
// NewSheet() will create a new additional sheet
func CreatExcel(path string) *excelize.File {
	excel := excelize.NewFile()

	// save excel file
	if err := excel.SaveAs(path); err != nil {
		log.Fatal(err)
		return nil
	}
	return excel
}

// Write table to excel Sheet1
// header should be a struct. How struct is based on network test purpose
func CreateTable(file string, header interface{}) {}
