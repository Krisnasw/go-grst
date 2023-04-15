package descriptor

type Descriptor struct {
	FileToGenerate []*FileDescriptorExt
	Repository     *DescriptorRepository
}

type DescriptorRepository struct {
	Files    map[string]*FileDescriptorExt
	Messages map[string]*MessageDescriptorExt
	Enums    map[string]*EnumDescriptorExt
}
