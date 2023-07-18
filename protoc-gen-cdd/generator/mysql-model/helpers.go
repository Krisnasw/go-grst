package mysql_model

import (
	"strings"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
)

func getGormTagAttribute(fieldext *descriptor.FieldDescriptorExt) string {
	result := ""
	if fieldext.MysqlField.PrimaryKey {
		result = "primary_key"
	}
	if fieldext.MysqlField.ColumnName != "" {
		if result != "" {
			result += ";"
		}
		result += "column:" + fieldext.MysqlField.ColumnName
	}
	if fieldext.MysqlField.ColumnType != "" {
		if result != "" {
			result += ";"
		}
		result += "type:" + fieldext.MysqlField.ColumnType
	}
	return result
}

func getGoType(fieldext *descriptor.FieldDescriptorExt) string {
	t := ""
	if fieldext.GetTypeName() == "" {
		t = strings.ToLower(strings.Replace(fieldext.GetType().String(), "TYPE_", "", -1))
		if t == "double" {
			t = "float64"
		}
	} else {
		switch fieldext.GetTypeName() {
		case ".google.protobuf.Timestamp":
			t = "time.Time"
		}
	}
	return t
}

func getGoStandartType(fieldext *descriptor.FieldDescriptorExt) string {
	t := getGoType(fieldext)
	switch t {
	case "float32":
		return "float64"
	case "float64":
		return t
	case "int":
		return t
	case "int8":
		return "int"
	case "int16":
		return "int"
	case "int32":
		return "int"
	case "int64":
		return "int"
	case "uint":
		return t
	case "uint16":
		return "uint"
	case "uint32":
		return "uint"
	case "uint64":
		return "uint"
	}
	return t
}
