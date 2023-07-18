package descriptor

import (
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

type ServiceDescriptorExt struct {
	*descriptorpb.ServiceDescriptorProto
	Repository *DescriptorRepository
	FileExt    *FileDescriptorExt
	MethodExt  []*MethodDescriptorExt
}

func (ServiceDescriptorExt) New(fext *FileDescriptorExt, svc *descriptorpb.ServiceDescriptorProto) *ServiceDescriptorExt {
	svcext := &ServiceDescriptorExt{
		ServiceDescriptorProto: svc,
		Repository:             fext.Repository,
		FileExt:                fext,
		MethodExt:              []*MethodDescriptorExt{},
	}
	for _, mth := range svcext.GetMethod() {
		mthext := MethodDescriptorExt{}.New(svcext, mth)
		svcext.MethodExt = append(svcext.MethodExt, mthext)
	}
	return svcext
}

func (svcext *ServiceDescriptorExt) GetIdentifier() string {
	return svcext.FileExt.GetIdentifier() + "." + svcext.GetName()
}
