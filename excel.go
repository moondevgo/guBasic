package guBasic

// https://programming.vip/docs/go-write-excel-file.html
// # Brief

import (
	"fmt"

	excelize "github.com/xuri/excelize/v2"
)

// # Type
// ## Excel(struct)
type Excel struct {
	Path  string         // file path
	Sheet string         // sheet name
	File  *excelize.File // file object
}

// ## Sheets(interface)
type Sheets interface {
	GetPath()
	GetSheet()
	SetPath() string
	SetSheet() string
	Open() *excelize.File
	// Close() nil
	Read() string
	Write() bool
}

// # Function(Private)
// ## indexOf
//   - data 내에서 element의 인덱스
func indexOf(element string, data []string) int {
	// fmt.Printf("%v in %v", element, data)
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// ## headerIndexes
//   - data 내에서 element의 인덱스
//
// header: []string{"h1", "h2", "h3", "h4", "h5"}, fields: []string{"h1", "h3", "h4"}
// indexes: map[string]int{{"h1": 0}, {"h3": 2}, {"h4": 3}}
func headerIndexes(header, fields []string) (indexes map[int]string) {
	indexes = make(map[int]string)
	for _, field := range fields {
		indexes[indexOf(field, header)] = field
	}
	return indexes
}

// ## dictFromSlices
//   - slices [][]string
//   - TODO: generic map[string]<T>
func dictsFromSlices(slices [][]string) (dicts []map[string]string) {
	dicts = []map[string]string{}

	indexes := headerIndexes(slices[0], slices[0])
	for _, row := range slices[1:] {
		dict := map[string]string{}
		for i, cell := range row {
			dict[indexes[i]] = cell
		}
		dicts = append(dicts, dict)
	}
	return dicts
}

// Excel 초기화
func InitExcel(path, sheet string) (e *Excel) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
	}
	e = &Excel{
		path, sheet, file,
	}

	return e
}

// Excel 생성
func NewExcel(path, sheet string) (e *Excel) {
	file := excelize.NewFile()
	e = &Excel{
		path, sheet, file,
	}

	return e
}

// # Function(interface implement)
// ## SetFile
//   - Set Excel's File
func (e *Excel) GetPath() string {
	return e.Path
}

func (e *Excel) GetSheet() string {
	return e.Sheet
}

func (e *Excel) GetFile() *excelize.File {
	return e.File
}

func (e *Excel) SetPath(filePath string) {
	e.Path = filePath
}

func (e *Excel) SetSheet(sheetName string) {
	e.Sheet = sheetName
}

func (e *Excel) SetFile(filePath string) (file *excelize.File) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func (e *Excel) Open() *excelize.File {
	f, err := excelize.OpenFile(e.Path)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

func (e *Excel) Read() []map[string]string {
	f, err := excelize.OpenFile(e.Path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	rows, err := f.GetRows(e.Sheet)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return dictsFromSlices(rows)
}

// Write By Cell
//
//	cells: [][2]string{
//			{"A2", "Hello World."},
//			{"B2", "100"},
//		}
func (e *Excel) WriteCells(sheetName string, cells [][2]string) bool {
	f := excelize.NewFile()
	index := f.NewSheet(sheetName)
	for _, cell := range cells {
		f.SetCellValue(sheetName, cell[0], cell[1]) //
	}
	f.SetActiveSheet(index)
	if err := f.SaveAs(e.Path); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Write By Rows
func (e *Excel) Write(sheetName string, rows []map[string]interface{}) bool {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName) //Set the name of the worksheet
	// header := ""

	for i, obj := range rows {
		//--Concatenate cell names based on rows and columns
		name, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to splice cell names,error:%s", err))
			return false
		}
		err = f.SetSheetRow(sheetName, name, &obj)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to write data by row,error:%s", err))
			return false
		}
	}

	if err := f.SaveAs(e.Path); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
