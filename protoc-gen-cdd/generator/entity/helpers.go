package entity

import (
	"fmt"
	"strings"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
)

func needImportTime(mext *descriptor.MessageDescriptorExt) bool {
	needImport := false
	if !needImport {
		for _, fext := range mext.FieldExt {
			if getGoType(fext) == "time.Time" {
				needImport = true
				break
			}
		}
	}
	return needImport
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
		default:

			t = strings.Replace(fieldext.GetTypeName(), fmt.Sprintf(".%s.", fieldext.MessageExt.FileExt.GetPackage()), "", -1)
		}
	}
	return t
}

func isRepeatTypeField(fieldext *descriptor.FieldDescriptorExt) bool {
	return fieldext.GetLabel().String() == "LABEL_REPEATED"
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
