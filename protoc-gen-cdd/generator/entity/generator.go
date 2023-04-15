package entity

import (
	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
)

type EntityGeneratorTemplate struct {
	name       string
	entityName []string
	descriptor *descriptor.Descriptor
}

func New(d *descriptor.Descriptor, entityName []string) *EntityGeneratorTemplate {
	result := &EntityGeneratorTemplate{
		name:       "entity",
		descriptor: d,
		entityName: entityName,
	}

	return result
}

func (t *EntityGeneratorTemplate) Generate() ([]*generator.GeneratorResponseFile, error) {
	mapOfEntityToGenerate := map[string]bool{}
	for _, e := range t.entityName {
		mapOfEntityToGenerate[e] = true
	}

	var files []*generator.GeneratorResponseFile
	for _, f := range t.descriptor.FileToGenerate {
		for _, mext := range f.MessageExt {
			if _, ok := mapOfEntityToGenerate[mext.GetName()]; ok {
				fileEntity, err := ApplyTemplateEntity(mext, f)
				if err != nil {
					return nil, err
				}

				files = append(files, fileEntity)
			}

		}
	}
	return files, nil
}
