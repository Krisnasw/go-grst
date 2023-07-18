package descriptor

import (
	"strings"

	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type EnumDescriptorExt struct {
	*descriptorpb.EnumDescriptorProto
	Repository *DescriptorRepository
	FileExt    *FileDescriptorExt
	NestedPath []string
}

func (EnumDescriptorExt) New(fext *FileDescriptorExt, enum *descriptorpb.EnumDescriptorProto, nestedPath []string) *EnumDescriptorExt {
	enumext := &EnumDescriptorExt{
		EnumDescriptorProto: enum,
		Repository:          fext.Repository,
		FileExt:             fext,
		NestedPath:          nestedPath,
	}
	return enumext
}

func (enumext *EnumDescriptorExt) GetIdentifier() string {
	components := []string{""}
	if enumext.FileExt.Package != nil {
		components = append(components, enumext.FileExt.GetPackage())
	}
	components = append(components, enumext.NestedPath...)
	components = append(components, enumext.GetName())
	return strings.Join(components, ".")
}
