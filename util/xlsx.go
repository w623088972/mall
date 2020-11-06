package util

import (
	"fmt"
	"reflect"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var columnIndex = make([]rune, 26)

type ColumnConverter func(v interface{}) interface{}
type XLSX struct {
	file       *excelize.File
	rowIdx     int
	sheetName  string
	fileName   string
	sheetIndex int
	converters map[string]ColumnConverter
}

func init() {
	for i := 0; i < 26; i++ {
		columnIndex[i] = 'A' + rune(i)
	}
}

func NewXLSX(fileName string) *XLSX {
	x := &XLSX{
		file:       excelize.NewFile(),
		rowIdx:     0,
		sheetName:  "Sheet1",
		fileName:   fileName,
		converters: make(map[string]ColumnConverter, 0),
	}
	x.sheetIndex = x.file.NewSheet(x.sheetName)
	return x
}

func (self *XLSX) AddColumnConverters(name string, converter ColumnConverter) {
	_, ok := self.converters[name]
	if ok {
		msg := fmt.Sprintf("column converter for column:%s already exists\n", name)
		panic(msg)
	}
	self.converters[name] = converter
}

func (self *XLSX) WriteTitle(v interface{}) {
	typeInfo := reflect.TypeOf(v).Elem()
	column := 0
	for k := 0; k < typeInfo.NumField(); k++ {
		field := typeInfo.Field(k)
		tag := field.Tag.Get("title")
		if tag == "" {
			continue
		}
		pos := fmt.Sprintf("%c%d", columnIndex[column%26], 1)
		self.file.SetCellValue(self.sheetName, pos, tag)
		column++
	}
}

func (self *XLSX) Write(v interface{}) {
	column := 0
	typeInfo := reflect.TypeOf(v).Elem()
	value := reflect.ValueOf(v).Elem()
	for k := 0; k < typeInfo.NumField(); k++ {
		field := typeInfo.Field(k)
		tag := field.Tag.Get("title")
		if tag == "" {
			continue
		}
		//fieldType := field.Type
		fieldV := value.Field(k)
		pos := fmt.Sprintf("%c%d", columnIndex[column%26], self.rowIdx+2)
		column++
		name := field.Name
		if handler, ok := self.converters[name]; ok {
			self.file.SetCellValue(self.sheetName, pos, handler(fieldV))
		} else {
			self.file.SetCellValue(self.sheetName, pos, fieldV)
		}
	}
	self.rowIdx++
}

func (self *XLSX) Save() error {
	// 设置工作簿的默认工作表
	self.file.SetActiveSheet(self.sheetIndex)
	// 根据指定路径保存文件
	//fileName := fmt.Sprintf("%s.xlsx", self.fileName)
	err := self.file.SaveAs(self.fileName)
	if err != nil {
		return err
	}
	return nil
}
