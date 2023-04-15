package descriptor

import (
	"fmt"
	"log"
	"strings"

	cddext "github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/cdd/api"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FieldDescriptorExt struct {
	*descriptorpb.FieldDescriptorProto
	Repository      *DescriptorRepository
	MessageExt      *MessageDescriptorExt
	MysqlField      *cddext.MysqlField
	ValidationRules []string
	DefaultValueExt string
}

func (FieldDescriptorExt) New(msgext *MessageDescriptorExt, field *descriptorpb.FieldDescriptorProto) *FieldDescriptorExt {
	fieldext := &FieldDescriptorExt{
		FieldDescriptorProto: field,
		Repository:           msgext.Repository,
		MessageExt:           msgext,
		MysqlField:           nil,
		ValidationRules:      []string{},
		DefaultValueExt:      "",
	}
	fieldext.MysqlField = parseExtMysqlField(field)
	if fieldext.MysqlField == nil {
		fieldext.MysqlField = &cddext.MysqlField{ColumnName: fieldext.GetJsonName(), PrimaryKey: false, ColumnType: ""}
	} else if fieldext.MysqlField.ColumnName == "" {
		fieldext.MysqlField.ColumnName = fieldext.GetJsonName()
	}

	validationRule := parseExtFieldValidation(field)
	if validationRule != "" {
		fieldext.ValidationRules = strings.Split(validationRule, "|")
	}

	fieldext.DefaultValueExt = parseExtFieldDefault(field)

	return fieldext
}

func (fieldext *FieldDescriptorExt) GetIdentifier() string {
	return fieldext.MessageExt.GetIdentifier() + "." + fieldext.GetName()
}

func parseExtMysqlField(field *descriptorpb.FieldDescriptorProto) *cddext.MysqlField {
	if field.Options == nil {
		return nil
	} else if !proto.HasExtension(field.Options, cddext.E_MysqlField) {
		return nil
	}

	ext := proto.GetExtension(field.Options, cddext.E_MysqlField)
	opts, ok := ext.(*cddext.MysqlField)
	if !ok {
		log.Println(fmt.Errorf("[parseExtMysqlField] extension is %T; want an MysqlField", ext))
		return nil
	}
	return opts
}

// func parseExtDBField(field *descriptorpb.FieldDescriptorProto) *cddext.DBField {
// 	if field.Options == nil {
// 		return nil
// 	} else if !proto.HasExtension(field.Options, cddext.E_Dbfield) {
// 		return nil
// 	}

// 	ext := proto.GetExtension(field.Options, cddext.E_Dbfield)
// 	opts, ok := ext.(*cddext.DBField)
// 	if !ok {
// 		log.Println(fmt.Errorf("[parseExtDBField] extension is %T; want an DBField", ext))
// 		return nil
// 	}
// 	return opts
// }

func parseExtFieldValidation(field *descriptorpb.FieldDescriptorProto) string {
	if field.Options == nil {
		return ""
	} else if !proto.HasExtension(field.Options, cddext.E_Validate) {
		return ""
	}

	ext := proto.GetExtension(field.Options, cddext.E_Validate)
	opts, ok := ext.(string)
	if !ok {
		log.Println(fmt.Errorf("[parseExtFieldValidation] extension is %T; want an string", ext))
		return ""
	}
	return opts
}

func parseExtFieldDefault(field *descriptorpb.FieldDescriptorProto) string {
	if field.Options == nil {
		return ""
	} else if !proto.HasExtension(field.Options, cddext.E_Default) {
		return ""
	}

	ext := proto.GetExtension(field.Options, cddext.E_Default)
	opts, ok := ext.(string)
	if !ok {
		log.Println(fmt.Errorf("[parseExtFieldDefault] extension is %T; want an string", ext))
		return ""
	}
	return opts
}
