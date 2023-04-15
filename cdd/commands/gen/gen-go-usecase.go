package gen

import (
	"fmt"

	protocgencdd "github.com/krisnasw/go-grst/cdd/cli/protoc-gen-cdd"
	usecase "github.com/krisnasw/go-grst/cdd/pkg/gen-go-usecase"
	"github.com/spf13/cobra"
)

type GenGoUsecase struct {
	Command         *cobra.Command
	protocGenCddCli *protocgencdd.ProtocGenCdd
	name            string
	output          string
	printLog        bool
}

func NewGenGoUsecaseCmd() *GenGoUsecase {
	c := &GenGoUsecase{
		Command: &cobra.Command{
			Use:     "go-usecase",
			Aliases: []string{"go-uc"},
			Short:   "generate cdd's empty usecase",
			Long:    "generate cdd's empty usecase",
			Run:     nil,
		},
		protocGenCddCli: protocgencdd.NewProtocGenCdd(),
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.name, "name", "n", "", "usecase name")
	c.Command.Flags().StringVarP(&c.output, "output", "o", "", "output directory. default = app/usecase")
	c.Command.Flags().BoolVar(&c.printLog, "print", false, "print log")
	return c
}

func (c *GenGoUsecase) runCommand(cmd *cobra.Command, args []string) error {
	if c.name == "" {
		return fmt.Errorf("--name or -n is required")
	}
	err := usecase.Generate(c.name, c.output)

	return err
}
