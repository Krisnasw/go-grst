package descriptor

import (
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type FileDescriptorExt struct {
	*descriptorpb.FileDescriptorProto
	Repository *DescriptorRepository
	ServiceExt []*ServiceDescriptorExt
	MessageExt []*MessageDescriptorExt
	EnumExt    []*EnumDescriptorExt
}

func (FileDescriptorExt) New(file *descriptorpb.FileDescriptorProto, repository *DescriptorRepository) *FileDescriptorExt {
	fext := &FileDescriptorExt{
		FileDescriptorProto: file,
		ServiceExt:          []*ServiceDescriptorExt{},
		MessageExt:          []*MessageDescriptorExt{},
		EnumExt:             []*EnumDescriptorExt{},
		Repository:          repository,
	}
	return fext
}

func (fext *FileDescriptorExt) GetIdentifier() string {
	return fext.FileDescriptorProto.GetName()
}
