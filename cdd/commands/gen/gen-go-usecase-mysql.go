package gen

import (
	"fmt"

	gocli "github.com/krisnasw/go-grst/cdd/cli/go"
	protocgencdd "github.com/krisnasw/go-grst/cdd/cli/protoc-gen-cdd"
	"github.com/spf13/cobra"
)

type GenGoUsecaseMysql struct {
	Command         *cobra.Command
	protocGenCddCli *protocgencdd.ProtocGenCdd
	protoPath       string
	protoFile       string
	name            string
	printLog        bool
}

func NewGenGoUsecaseMysqlCmd() *GenGoUsecaseMysql {
	c := &GenGoUsecaseMysql{
		Command: &cobra.Command{
			Use:     "go-usecase-mysql",
			Aliases: []string{"go-ucm"},
			Short:   "generate cdd's crud mysql usecase",
			Long:    "generate cdd's crud mysql usecase",
			Run:     nil,
		},
		protocGenCddCli: protocgencdd.NewProtocGenCdd(),
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.protoPath, "proto-path", "p", "contract/", "contract proto filepath")
	c.Command.Flags().StringVarP(&c.protoFile, "proto-file", "f", "", "contract proto filename")
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "mysql model name (message type in *.proto")
	c.Command.Flags().BoolVar(&c.printLog, "print", false, "print log")
	return c
}

func (c *GenGoUsecaseMysql) runCommand(cmd *cobra.Command, args []string) error {
	if c.protoFile == "" {
		return fmt.Errorf("--proto-file or -f is required")
	} else if c.name == "" {
		return fmt.Errorf("--name or -n is required")
	}

	goModuleName, err := gocli.GetCurrentModule()
	if err != nil {
		return err
	}
	err = c.protocGenCddCli.GenerateUsecaseMysql(c.protoFile, c.protoPath, "", c.name, goModuleName, c.printLog)

	return err
}
