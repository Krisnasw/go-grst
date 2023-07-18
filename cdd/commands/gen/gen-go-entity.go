package gen

import (
	"fmt"
	"strings"

	protocgencdd "github.com/krisnasw/go-grst/cdd/cli/protoc-gen-cdd"
	"github.com/spf13/cobra"
)

type GenGoEntity struct {
	Command         *cobra.Command
	protocGenCddCli *protocgencdd.ProtocGenCdd
	entities        string
	protoPath       string
	protoFile       string
	outputPath      string
	printLog        bool
}

func NewGenGoEntityCmd() *GenGoEntity {
	c := &GenGoEntity{
		Command: &cobra.Command{
			Use:     "go-entity",
			Aliases: []string{"go-e"},
			Short:   "generate cdd's entity",
			Long:    "generate cdd's entity",
			Run:     nil,
		},
		protocGenCddCli: protocgencdd.NewProtocGenCdd(),
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.protoPath, "proto-path", "p", "contract/", "contract proto filepath")
	c.Command.Flags().StringVarP(&c.protoFile, "proto-file", "f", "", "contract proto filename")
	c.Command.Flags().StringVarP(&c.entities, "entity", "e", "", "list of entity to generate (based on proto message). ex: Entity1|Entity2")
	c.Command.Flags().StringVarP(&c.outputPath, "output", "o", "entity/", "output path for generated file. default: entity/*")
	c.Command.Flags().BoolVar(&c.printLog, "print", false, "print log")
	return c
}

func (c *GenGoEntity) runCommand(cmd *cobra.Command, args []string) error {
	if c.protoFile == "" {
		return fmt.Errorf("--proto-file or -f is required")
	} else if c.entities == "" {
		return fmt.Errorf("--entity or -e is required")
	}

	err := c.protocGenCddCli.GenerateEntity(c.protoFile, c.protoPath, c.outputPath, strings.Split(c.entities, "|"), c.printLog)

	return err
}
