package descriptor

import (
	"fmt"
	"log"
	"strings"

	cddext "github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/cdd/api"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type MessageDescriptorExt struct {
	*descriptorpb.DescriptorProto
	Repository *DescriptorRepository
	FileExt    *FileDescriptorExt
	NestedPath []string
	FieldExt   []*FieldDescriptorExt
	Mysql      *cddext.Mysql
}

func (MessageDescriptorExt) New(fext *FileDescriptorExt, msg *descriptorpb.DescriptorProto, nestedPath []string) *MessageDescriptorExt {
	msgext := &MessageDescriptorExt{
		DescriptorProto: msg,
		Repository:      fext.Repository,
		FileExt:         fext,
		NestedPath:      nestedPath,
		FieldExt:        []*FieldDescriptorExt{},
		Mysql:           nil,
	}
	for _, field := range msgext.Field {
		fieldext := FieldDescriptorExt{}.New(msgext, field)
		msgext.FieldExt = append(msgext.FieldExt, fieldext)
	}

	msgext.Mysql = parseExtMysql(msg)
	if msgext.Mysql == nil {
		msgext.Mysql = &cddext.Mysql{TableName: strings.Replace(strcase.ToKebab(msgext.GetName()), "-", "_", -1), EnableSoftDelete: false, DisableTimestampTracking: true}
	}

	return msgext
}

func (msgext *MessageDescriptorExt) GetIdentifier() string {
	components := []string{""}
	if msgext.FileExt.Package != nil {
		components = append(components, msgext.FileExt.GetPackage())
	}
	components = append(components, msgext.NestedPath...)
	components = append(components, msgext.GetName())
	return strings.Join(components, ".")
}

func parseExtMysql(message *descriptorpb.DescriptorProto) *cddext.Mysql {
	if message.Options == nil {
		return nil
	} else if !proto.HasExtension(message.Options, cddext.E_Mysql) {
		return nil
	}

	ext := proto.GetExtension(message.Options, cddext.E_Mysql)
	opts, ok := ext.(*cddext.Mysql)
	if !ok {
		log.Println(fmt.Errorf("[parseExtMysql] extension is %T; want an Mysql", ext))
		return nil
	}
	return opts
}
