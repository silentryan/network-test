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
// type ExlHandler struct {
// 	sheet   string
// 	table   string
// 	headers []string
// }

// Get data of specific columns by row index
func GetContentByCols(fname string, sheet string, cols []string) [][]string {
	var values [][]string = make([][]string, 1000)

	// get file
	f, err := excelize.OpenFile(fname)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	// according to a sheet. get all rows
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Get cols index
	colsIndex := TransColsIndextoNum(cols)

	// loop rows and get filter specific cols
	// except the first row (header)
	for i := 1; i < len(rows); i++ {
		values[i] = make([]string, 0, 100)
		row := rows[i]
		for _, colIndex := range colsIndex {
			// 检查列索引是否在行的范围内，避免索引越界
			if colIndex < len(row) {
				values[i] = append(values[i], row[colIndex])
			} else {
				values[i] = append(values[i], "") // 超出范围则填充空字符串
			}
		}
	}
	return values
}

func TransColsIndextoNum(cols []string) []int {
	var colsIndex []int
	for _, col := range cols {
		// 将Excel列名（如"A"）转换为数字索引（1, 2, 3...）
		idx, err := excelize.ColumnNameToNumber(col)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		// col's index start from 0
		colsIndex = append(colsIndex, idx-1)
	}
	return colsIndex
}

// excelize can only write data into specific cell
func WriteTableRow(f *excelize.File, sheet string, rowIndex int, rowData []string) {
	for colIndex, coldata := range rowData {
		// 1. Get col Id from col index
		colId, err := excelize.ColumnNumberToName(colIndex + 1)
		if err != nil {
			fmt.Println("列索引转换失败:", err)
			return
		}

		// 2. Get cell index
		cellCoord := fmt.Sprintf("%s%d", colId, rowIndex+1)

		// 3. Write data
		if err := f.SetCellValue(sheet, cellCoord, coldata); err != nil {
			fmt.Println("写入单元格失败:", err)
			return
		}
	}
}

func WriteDatetoTable(fname string, sheet string, data [][]string) {
	// 1. 打开已存在的Excel文件
	f, err := excelize.OpenFile("你的文件.xlsx")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer f.Close()

	// 2.获取当前表格的总行数，以确定新行的位置
	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println("获取行数据失败:", err)
		return
	}
	currentRowIndex := len(rows)

	// 3. Write data by rows
	for rowIndex, row := range rows {
		WriteTableRow(f, sheet, currentRowIndex+rowIndex, row)
	}
}
