package gengousecase

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	tmplUseCaseRepoImpl = template.Must(template.New("usecase-repo-impl").Funcs(template.FuncMap{
		"ToSnake": strcase.ToSnake,
	}).Parse(`
	package {{ToSnake .UsecaseName}}
	
	type repository struct {
		//db     *gorm.DB
		//ds *{{ToSnake .UsecaseName}}_ds.MysqlDatasource
	}
	func NewRepository() Repository {
		return &repository{}
	}

	`))
)

func applyTemplateUseCaseRepoImpl(usecaseName string, filepath string) (*generatorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		UsecaseName string
	}{
		usecaseName,
	}

	if err := tmplUseCaseRepoImpl.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}
	return &generatorResponseFile{
		outputPath: filepath + "/repository-impl.cdd.go",
		content:    string(formatted),
	}, nil
}
