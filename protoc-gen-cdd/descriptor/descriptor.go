package descriptor

import (
	"log"

	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func New(r pluginpb.CodeGeneratorRequest) *Descriptor {
	d := &Descriptor{
		FileToGenerate: []*FileDescriptorExt{},
		Repository: &DescriptorRepository{
			Files:    map[string]*FileDescriptorExt{},
			Messages: map[string]*MessageDescriptorExt{},
			Enums:    map[string]*EnumDescriptorExt{},
		},
	}

	// Register to repository
	d.registerRepository(r)

	// set file to generate
	for _, filename := range r.FileToGenerate {
		if fext, ok := d.Repository.Files[filename]; ok {
			d.FileToGenerate = append(d.FileToGenerate, fext)
		} else {
			log.Println("[Warning] File `" + filename + "` not found on descriptor repository")
			continue
		}
	}
	return d

}
func (d *Descriptor) registerRepository(r pluginpb.CodeGeneratorRequest) {
	for _, file := range r.GetProtoFile() {
		fext := d.registerFile(file)
		d.registerMessage(fext, file.GetMessageType(), []string{})
		d.registerEnum(fext, file.GetEnumType(), []string{})
	}
	// service can be registered if messages already registered
	for _, fext := range d.Repository.Files {
		for _, svc := range fext.GetService() {
			svcext := ServiceDescriptorExt{}.New(fext, svc)
			fext.ServiceExt = append(fext.ServiceExt, svcext)
		}
	}
}

func (d *Descriptor) registerFile(file *descriptorpb.FileDescriptorProto) *FileDescriptorExt {
	fext := FileDescriptorExt{}.New(file, d.Repository)
	d.Repository.Files[fext.GetIdentifier()] = fext
	return fext
}

func (d *Descriptor) registerMessage(fext *FileDescriptorExt, msgs []*descriptorpb.DescriptorProto, nestedPath []string) {
	for _, msg := range msgs {
		mext := MessageDescriptorExt{}.New(fext, msg, nestedPath)
		fext.MessageExt = append(fext.MessageExt, mext)
		d.Repository.Messages[mext.GetIdentifier()] = mext

		d.registerMessage(fext, msg.GetNestedType(), append([]string{}, nestedPath...))
		d.registerEnum(fext, msg.GetEnumType(), append([]string{}, nestedPath...))
	}
}

func (d *Descriptor) registerEnum(fext *FileDescriptorExt, enums []*descriptorpb.EnumDescriptorProto, nestedPath []string) {
	for _, enum := range enums {
		enumext := EnumDescriptorExt{}.New(fext, enum, nestedPath)
		fext.EnumExt = append(fext.EnumExt, enumext)
		d.Repository.Enums[enumext.GetIdentifier()] = enumext
	}
}
