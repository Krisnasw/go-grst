package gen

import (
	"fmt"
	"path/filepath"

	mysql_query_gen "github.com/krisnasw/go-grst/cdd/pkg/gen-go-mysql-query"
	"github.com/spf13/cobra"
)

type GenGoDsMysql struct {
	Command   *cobra.Command
	inputFile string
	modelName string
	outputDir string
}

func NewGenGoDatasourceMysqlCmd() *GenGoDsMysql {
	c := &GenGoDsMysql{
		Command: &cobra.Command{
			Use:     "go-mysql-query",
			Aliases: []string{"go-msq"},
			Short:   "generate query mysql query based on struct model",
			Long:    "generate query mysql query based on struct model",
			Run:     nil,
		},
	}
	c.Command.RunE = c.runCommand
	c.Command.Flags().StringVarP(&c.inputFile, "input", "f", "", "input file path (*.go file)")
	c.Command.Flags().StringVarP(&c.modelName, "model-name", "n", "", "struct model name (case insensitive)")
	c.Command.Flags().StringVarP(&c.outputDir, "output-dir", "o", "", "output directory (optional). Default: same with input file's directory")
	return c
}

func (c *GenGoDsMysql) runCommand(cmd *cobra.Command, args []string) error {
	if c.inputFile == "" {
		return fmt.Errorf("--input or -f is required")
	}
	if c.outputDir == "" {
		c.outputDir = filepath.Dir(c.inputFile)
	}

	err := mysql_query_gen.Generate(c.inputFile, c.modelName, c.outputDir)
	return err
}
