package usecase_mysql

import (
	"fmt"
	"log"

	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
)

type ScaffoldMysqlGeneratorTemplate struct {
	name         string
	modelName    string
	goModuleName string
	descriptor   *descriptor.Descriptor
}

func New(d *descriptor.Descriptor, modelName string, goModuleName string) *ScaffoldMysqlGeneratorTemplate {
	result := &ScaffoldMysqlGeneratorTemplate{
		name:         "usecase-mysql",
		modelName:    modelName,
		descriptor:   d,
		goModuleName: goModuleName,
	}

	return result
}

func (t *ScaffoldMysqlGeneratorTemplate) Generate() ([]*generator.GeneratorResponseFile, error) {
	var files []*generator.GeneratorResponseFile
	for _, f := range t.descriptor.FileToGenerate {
		for _, mext := range f.MessageExt {
			if mext.GetName() != t.modelName {
				continue
			} else if !mext.Mysql.DbModel {
				log.Println(fmt.Sprintf("[Warning] `%s` was found, but db_model: false", t.modelName))
				continue
			}

			fileUseCaseErrors, err := applyTemplateUseCaseErrors(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseErrors)

			fileUseCaseIntf, err := applyTemplateUseCaseIntf(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseIntf)

			fileUseCaseRepoImpl, err := applyTemplateUseCaseRepoImpl(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseRepoImpl)

			fileUseCaseImpl, err := applyTemplateUseCaseImpl(ScaffoldMysql{mext, t.goModuleName})
			if err != nil {
				return nil, err
			}
			files = append(files, fileUseCaseImpl)
		}
	}

	return files, nil
}
