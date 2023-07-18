package generator

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Generator interface {
	Generate() ([]*GeneratorResponseFile, error)
}

type GeneratorResponseFile struct {
	Filename     string
	GoImportPath protogen.GoImportPath
	Content      string
}
