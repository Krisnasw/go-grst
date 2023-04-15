package gen

import (
	"github.com/spf13/cobra"
)

type GenCmd struct {
	*cobra.Command
}

func NewGenCmd() *GenCmd {
	c := &GenCmd{}
	c.Command = &cobra.Command{
		Use:   "gen",
		Short: "Generate cdd: grst framework + crud",
		Long:  "Generate contract driven development (cdd) grst & crud template based on service.yaml through protoc-gen-cdd",
	}
	c.AddCommand(NewGenGoCmd().Command)
	c.AddCommand(NewGenGoEntityCmd().Command)
	c.AddCommand(NewGenGoUsecaseCmd().Command)
	c.AddCommand(NewGenGoUsecaseMysqlCmd().Command)
	c.AddCommand(NewGenGoDatasourceMysqlCmd().Command)
	return c
}
