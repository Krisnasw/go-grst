package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/golang/glog"
	descriptor "github.com/krisnasw/go-grst/protoc-gen-cdd/descriptor"
	"github.com/krisnasw/go-grst/protoc-gen-cdd/generator"
	entity "github.com/krisnasw/go-grst/protoc-gen-cdd/generator/entity"
	grstframework "github.com/krisnasw/go-grst/protoc-gen-cdd/generator/grst-framework"
	mysql_model "github.com/krisnasw/go-grst/protoc-gen-cdd/generator/mysql-model"
	usecase_mysql "github.com/krisnasw/go-grst/protoc-gen-cdd/generator/usecase-mysql"

	"google.golang.org/protobuf/compiler/protogen"
)

var listOfType = "grst|scaffold-mysql|entity|usecase"
var (
	fType = flag.String("type", "grst", "option: "+listOfType)
	/*grst specific options*/
	fProtocGoOut = flag.Bool("protoc-gen-go", true, "generate *.pb.go (calling `protoc-gen-go`) with additional features, such as request validation & default value. protoc-gen-go version: v1.25.0. default: true")
	/*scaffold-mysql specific options*/
	fGoModuleName = flag.String("go-module-name", "", "Go module name, check in go.mod file. This needed for local import prefix. example: github.com/krisnasw/go-grst/examples/province-api")

	fName = flag.String("name", "", "name of entity or usecase or model")

	fVersion = flag.Bool("version", false, "version")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	if *fVersion {
		fmt.Println("protoc-gen-cdd v1.0.0")
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(func(plugin *protogen.Plugin) error {
		registry := descriptor.New(*plugin.Request)
		var gen generator.Generator = nil
		switch *fType {
		case "grst":
			gen = grstframework.New(registry, *plugin, *fProtocGoOut)
		case "mysql-model":
			gen = mysql_model.New(registry)
		case "entity":
			if *fName == "" {
				return fmt.Errorf("Option `entity` is required. Example `--cdd_opt entity=\"EntityName1|EntityName2\">`")
			}
			gen = entity.New(registry, strings.Split(*fName, "|"))
		case "usecase-mysql":
			if *fGoModuleName == "" {
				return fmt.Errorf("Option `go-module-name` is required. Example `--cdd_opt go-module-name=$(go list -m)>`")
			} else if *fName == "" {
				return fmt.Errorf("Option `name` is required. Example `--cdd_opt name=ModelName>`")
			}
			gen = usecase_mysql.New(registry, *fName, *fGoModuleName)

		default:
			return fmt.Errorf("Invalid option `type`, got: %s, expect: %s", *fType, listOfType)
		}
		if gen != nil {
			files, err := gen.Generate()
			if err != nil {
				return err
			}
			for _, f := range files {

				genFile := plugin.NewGeneratedFile(f.Filename, f.GoImportPath)
				if _, err := genFile.Write([]byte(f.Content)); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
