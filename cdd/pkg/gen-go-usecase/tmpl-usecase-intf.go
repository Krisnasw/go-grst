package gengousecase

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	tmplUseCaseIntf = template.Must(template.New("usecase-intf").Funcs(template.FuncMap{
		"ToSnake": strcase.ToSnake,
	}).Parse(`
	package {{ToSnake .UsecaseName}}
	
	type UseCase interface {
		// Get(id int) (*entity.{{.UsecaseName }}, error)
		// GetAll() ([]*entity.{{.UsecaseName }}, error)
		// Create(in entity.{{.UsecaseName}}) (*entity.{{.UsecaseName}}, error)
		// Update(in entity.{{.UsecaseName}}) (*entity.{{.UsecaseName}}, error)
		// Delete(id int) error
	}
	

	`))
)

func applyTemplateUseCaseIntf(usecaseName string, filepath string) (*generatorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		UsecaseName string
	}{
		usecaseName,
	}

	if err := tmplUseCaseIntf.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}
	return &generatorResponseFile{
		outputPath: filepath + "/interface.cdd.go",
		content:    string(formatted),
	}, nil
}
