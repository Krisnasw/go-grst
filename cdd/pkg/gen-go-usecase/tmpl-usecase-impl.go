package gengousecase

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	tmplUseCaseImpl = template.Must(template.New("usecase-impl").Funcs(template.FuncMap{
		"ToSnake": strcase.ToSnake,
	}).Parse(`
	package {{ToSnake .UsecaseName}}
	
	type usecase struct {
		// add repository here
	}
	
	func NewUsecase() UseCase {
		return &usecase{}
	}
	// func (uc *usecase) Get(id int) (*entity.{{.UsecaseName }}, error){}
	`))
)

func applyTemplateUseCaseImpl(usecaseName string, filepath string) (*generatorResponseFile, error) {
	w := bytes.NewBuffer(nil)

	var tmplData = struct {
		UsecaseName string
	}{
		usecaseName,
	}

	if err := tmplUseCaseImpl.Execute(w, tmplData); err != nil {
		return nil, err
	}

	formatted, err := format.Source([]byte(w.String()))
	if err != nil {
		return nil, err
	}

	return &generatorResponseFile{
		outputPath: filepath + "/implement.cdd.go",
		content:    string(formatted),
	}, nil
}
