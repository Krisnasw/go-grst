package protocgencdd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/krisnasw/go-grst/cdd/cli/protoc"
)

type ProtocGenCdd struct {
}

func NewProtocGenCdd() *ProtocGenCdd {
	return &ProtocGenCdd{}
}

func (pgc *ProtocGenCdd) GenerateGrst(protoFilename string, inputPath string, outputPath string, printLog bool) error {
	outputPath = outputPath + strings.Replace(filepath.Base(protoFilename), filepath.Ext(protoFilename), "", -1)
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	// p.AddProtoPath("$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "grst"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "go-grpc", Opts: map[string]string{}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "grpc-gateway", Opts: map[string]string{"logtostderr": "true", "generate_unbound_methods": "true"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=grst]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateMysqlModel(protoFilename string, inputPath string, outputPath string, printLog bool) error {
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "mysql-model"}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=mysql-model]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateEntity(protoFilename string, inputPath string, outputPath string, entities []string, printLog bool) error {
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "entity", "name": strings.Join(entities, "|")}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=entity]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}

func (pgc *ProtocGenCdd) GenerateUsecaseMysql(protoFilename string, inputPath string, outputPath string, modelName string, goModuleName string, printLog bool) error {
	os.MkdirAll(outputPath, os.ModePerm)

	p := protoc.NewProtoc()
	p.AddProtoPath(inputPath)
	p.AddProtoPath("$GOPATH/src/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/cddapis/")
	p.AddProtoPath("$GOPATH/src/github.com/krisnasw/go-grst/protoc-gen-cdd/ext/googleapis/")
	p.AddProtocGenOut(protoc.ProtocGenOut{Name: "cdd", Opts: map[string]string{"type": "usecase-mysql", "name": modelName, "go-module-name": goModuleName}, OutputPath: outputPath, Version: protoc.ProtobufVersion2})

	log.Println("Generating file [type=usecase-mysql]: " + inputPath + protoFilename + " | outpath: ./" + outputPath)
	err := p.Exec(filepath.Base(protoFilename), printLog)
	if err != nil {
		return err
	}
	return nil
}
