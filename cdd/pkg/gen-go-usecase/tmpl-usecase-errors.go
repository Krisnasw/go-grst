package gengousecase

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	tmplUseCaseErrors = template.Must(template.New("usecase-errors").Funcs(template.FuncMap{
		"ToSnake": strcase.ToSnake,
	}).Parse(`
	package {{ToSnake .UsecaseName}}
	
	import "errors"

	var ErrDatabaseError = errors.New("Database Error")
	var ErrRecordNotFound = errors.New("Record Not Found")

	`))
)

func applyTemplateUseCaseErrors(usecaseName string, filepath string) (*generatorResponseFile, error) {

	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		UsecaseName string
	}{
		usecaseName,
	}

	if err := tmplUseCaseErrors.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}
	return &generatorResponseFile{
		outputPath: filepath + "/errors.cdd.go",
		content:    string(formatted),
	}, nil

}
