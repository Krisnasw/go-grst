package grstframework

import (
	"fmt"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

type GrstFrameworkTemplate struct {
	name         string
	plugin       protogen.Plugin
	descriptor   *descriptor.Descriptor
	generatePbGo bool
}

func New(d *descriptor.Descriptor, plugin protogen.Plugin, generatePbGo bool) *GrstFrameworkTemplate {
	result := &GrstFrameworkTemplate{
		name:         "grst-framework",
		plugin:       plugin,
		descriptor:   d,
		generatePbGo: generatePbGo,
	}

	return result
}

func (t *GrstFrameworkTemplate) Generate() ([]*generator.GeneratorResponseFile, error) {
	var files []*generator.GeneratorResponseFile
	for _, f := range t.descriptor.FileToGenerate {
		fileGrpcRest, err := applyTemplateGrpcRest(f)
		if err != nil {
			return nil, err
		}
		files = append(files, fileGrpcRest)
	}

	/* Generate Protoc Gen Go */
	if t.generatePbGo {
		t.plugin.SupportedFeatures = gengo.SupportedFeatures
		for _, file := range t.plugin.Files {
			if file.Generate {
				genfile := gengo.GenerateFile(&t.plugin, file)
				genfile.Skip()

				fext, ok := t.descriptor.Repository.Files[file.Proto.GetName()]
				if !ok {
					return nil, fmt.Errorf(file.Proto.GetName() + ": not found in descriptor files repository")
				}
				fileGenGo, err := callProtocGoOut(fext, genfile, file)
				if err != nil {
					return nil, err
				}
				files = append(files, fileGenGo)
			}
		}
	}

	return files, nil
}
