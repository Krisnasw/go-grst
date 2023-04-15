package mysql_model

import (
	"github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
)

type MysqlModelGeneratorTemplate struct {
	name         string
	goModuleName string
	descriptor   *descriptor.Descriptor
}

func New(d *descriptor.Descriptor) *MysqlModelGeneratorTemplate {
	result := &MysqlModelGeneratorTemplate{
		name:       "mysql-model",
		descriptor: d,
	}

	return result
}

func (t *MysqlModelGeneratorTemplate) Generate() ([]*generator.GeneratorResponseFile, error) {
	var files []*generator.GeneratorResponseFile
	for _, f := range t.descriptor.FileToGenerate {
		for _, mext := range f.MessageExt {
			if !mext.Mysql.DbModel {
				continue
			}
			// fileEntity, err := entity.ApplyTemplateEntity(mext, f)
			// if err != nil {
			// 	return nil, err
			// }
			// files = append(files, fileEntity)

			fileRepoMysql, err := applyTemplateRepoMysql(MysqlModelParam{mext})
			if err != nil {
				return nil, err
			}
			files = append(files, fileRepoMysql)

			fileMysqlModel, err := applyTemplateMysqlModel(MysqlModelParam{mext}, f)
			if err != nil {
				return nil, err
			}
			files = append(files, fileMysqlModel)

		}
	}

	return files, nil
}
