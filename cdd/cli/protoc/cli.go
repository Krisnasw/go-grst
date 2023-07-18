package protoc

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func NewProtoc() Protoc {
	return &protoc{protoPaths: []string{}, protocGenOuts: []ProtocGenOut{}}
}

func (cli *protoc) AddProtoPath(pp ...string) {
	cli.protoPaths = append(cli.protoPaths, pp...)
}

func (cli *protoc) AddProtocGenOut(pgo ...ProtocGenOut) {
	cli.protocGenOuts = append(cli.protocGenOuts, pgo...)
}

func (cli *protoc) Exec(input string, printExecStatement bool) error {
	args := []string{}
	for _, p := range cli.protoPaths {
		args = append(args, fmt.Sprintf("--proto_path=%s", os.ExpandEnv(p)))
	}

	for _, out := range cli.protocGenOuts {
		if out.Version == ProtobufVersion1 {
			p := ""
			if len(out.Opts) == 0 {
				p = out.OutputPath
			} else {
				tmp := []string{}
				for k, v := range out.Opts {
					tmp = append(tmp, fmt.Sprintf("%s=%s", k, v))
				}
				p = fmt.Sprintf("%s:./%s", strings.Join(tmp, ","), out.OutputPath)
			}
			args = append(args, fmt.Sprintf("--%s_out=%s", out.Name, p))
		} else {
			args = append(args, fmt.Sprintf("--%s_out", out.Name), fmt.Sprintf("./%s", out.OutputPath))
			for k, v := range out.Opts {
				args = append(args, fmt.Sprintf("--%s_opt", out.Name), fmt.Sprintf("%s=%s", k, v))
			}

		}
	}
	args = append(args, input)
	terminalCmd := exec.Command("protoc", args...)
	if printExecStatement {
		log.Println("protoc", args, "\n")
	}
	terminalCmd.Stdin, terminalCmd.Stdout, terminalCmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	return terminalCmd.Run()
}
