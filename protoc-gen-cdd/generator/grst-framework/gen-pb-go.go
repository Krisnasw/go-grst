package grstframework

import (
	"fmt"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/gomodifytags"
	"google.golang.org/protobuf/compiler/protogen"
)

func callProtocGoOut(f *descriptor.FileDescriptorExt, genfile *protogen.GeneratedFile, protogenFile *protogen.File) (*generator.GeneratorResponseFile, error) {
	content, err := genfile.Content()
	if err != nil {
		return nil, err
	}
	parsedContent := string(content)

	for _, mext := range f.MessageExt {
		if mext == nil {
			continue
		}
		structName := mext.GetName()
		for _, fext := range mext.FieldExt {
			if fext == nil {
				continue
			}
			structNameParam := structName
			if fext.OneofIndex != nil {
				structNameParam += "_" + fext.GetName()
				// log.Println(structName, fext.GetName(), fext.GetJsonName())
				// continue
			}
			var err error
			fieldName := fext.GetName()
			parsedContent, err = gomodifytags.OverrideJSON_Content(parsedContent, structNameParam, fieldName, fext.GetJsonName(), true)
			if err != nil {
				return nil, fmt.Errorf("[error] gomodifytags.OverrideJSON_Content %s %s %s. Error: %s", structName, fieldName, fext.GetJsonName(), err.Error())
			}

			if len(fext.ValidationRules) > 0 {
				parsedContent, err = gomodifytags.AddValidate_Content(parsedContent, structName, fieldName, fext.ValidationRules)
				if err != nil {
					return nil, fmt.Errorf("[error] gomodifytags.AddValidate_Content %s %s %v. Error: %s", structName, fieldName, fext.ValidationRules, err.Error())
				}
			}

			if fext.DefaultValueExt != "" {
				parsedContent, err = gomodifytags.AddDefault_Content(parsedContent, structName, fieldName, fext.DefaultValueExt)
				if err != nil {
					return nil, fmt.Errorf("[error] gomodifytags.AddDefault_Content %s %s %v. Error: %s", structName, fieldName, fext.ValidationRules, err.Error())
				}
			}
		}
	}

	return &generator.GeneratorResponseFile{
		Filename:     protogenFile.GeneratedFilenamePrefix + ".pb.go",
		GoImportPath: protogenFile.GoImportPath,
		Content:      parsedContent,
	}, nil
}
