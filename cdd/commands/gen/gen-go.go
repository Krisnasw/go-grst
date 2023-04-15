package gen

import (
	"path/filepath"

	protocgencdd "github.com/krisnasw/go-grst/cdd/cli/protoc-gen-cdd"
	"github.com/krisnasw/go-grst/cdd/pkg/serviceYaml"
	"github.com/spf13/cobra"
)

type GenGoCmd struct {
	Command         *cobra.Command
	protocGenCddCli *protocgencdd.ProtocGenCdd
	serviceYamlFile string
	printLog        bool
}

func NewGenGoCmd() *GenGoCmd {
	c := &GenGoCmd{
		Command: &cobra.Command{
			Use:   "go",
			Short: "generate cdd framework",
			Long:  "generate cdd framework",
			Run:   nil,
		},
		protocGenCddCli: protocgencdd.NewProtocGenCdd(),
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVar(&c.serviceYamlFile, "service-yaml", "service.yaml", "service.yaml file path")
	c.Command.Flags().BoolVar(&c.printLog, "print", false, "print log")
	return c
}

type ContractToGenerate struct {
	protoInput    string
	outputGrstDir string
}

const defaultContractOutputGrst = "handler/grst/"
const defaultDependencyOutputGrst = "clients/grst/"

func (c *GenGoCmd) runCommand(cmd *cobra.Command, args []string) error {
	svcYaml, err := serviceYaml.GetServiceYAML(c.serviceYamlFile)
	if err != nil {
		return err
	}

	contractOutputGrst := svcYaml.Contract.OutputGrst
	dependencyOutputGrst := svcYaml.Dependency.OutputGrst

	if contractOutputGrst == "" {
		contractOutputGrst = defaultContractOutputGrst
	}
	if dependencyOutputGrst == "" {
		dependencyOutputGrst = defaultDependencyOutputGrst
	}

	contractsToGenerate := []ContractToGenerate{}
	// Setup proto contract for main service
	for _, file := range svcYaml.Contract.ProtoFiles {
		contractsToGenerate = append(contractsToGenerate, ContractToGenerate{
			protoInput:    file,
			outputGrstDir: contractOutputGrst,
		})
	}

	// Setup proto contract for dependencies services
	for _, svcFilePath := range svcYaml.Dependency.Services {
		svcYamlDependency, err := serviceYaml.GetServiceYAML(svcFilePath)
		if err != nil {
			return err
		}

		dirDependency, _ := filepath.Split(svcFilePath)

		for _, file := range svcYamlDependency.Contract.ProtoFiles {
			contractsToGenerate = append(contractsToGenerate, ContractToGenerate{
				protoInput:    dirDependency + "/" + file,
				outputGrstDir: dependencyOutputGrst,
			})
		}
	}
	// generate grpc pb
	for _, ctg := range contractsToGenerate {
		dir, filename := filepath.Split(ctg.protoInput)
		err = c.protocGenCddCli.GenerateGrst(filename, dir, ctg.outputGrstDir, c.printLog)
		if err != nil {
			return err
		}

		// Generate mysql datasource from contract will deprecated
		// if ctg.mysqlModel {
		// 	err = c.protocGenCddCli.GenerateMysqlModel(filename, dir, ctg.outputMysqlModelDir, c.printLog)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	return nil
}
