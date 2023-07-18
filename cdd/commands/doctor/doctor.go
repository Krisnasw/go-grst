package doctor

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gosuri/uiprogress/util/strutil"
	"github.com/spf13/cobra"
)

type DoctorCmd struct {
	*cobra.Command
}

func NewDoctorCmd() *DoctorCmd {
	c := &DoctorCmd{}
	c.Command = &cobra.Command{
		Use:   "doctor",
		Short: "Check requirement to use this cli",
		Long:  "Check requirement to use this cli",
	}
	c.Command.RunE = c.runCommand
	return c
}

func (c *DoctorCmd) runCommand(cmd *cobra.Command, args []string) error {
	const (
		InfoColor    = "\033[1;34m%s\033[0m"
		NoticeColor  = "\033[1;36m%s\033[0m"
		WarningColor = "\033[1;33m%s\033[0m"
		ErrorColor   = "\033[1;31m%s\033[0m"
		DebugColor   = "\033[0;36m%s\033[0m"
		SuccessColor = "\033[1;32m%s\033[0m"
	)

	protocVersion, err := exec.Command(`protoc`, "--version").Output()
	if err == nil {
		fmt.Printf(SuccessColor, "[v] ")
		fmt.Println(fmt.Sprintf("protoc (%s)", fmt.Sprint(strings.Replace(string(protocVersion), "\n", "", -1))))
	} else {
		fmt.Printf(ErrorColor, "[x] ")
		fmt.Println("protoc")
		fmt.Printf(ErrorColor+"\n", fmt.Sprint(strutil.Resize("", 4), "protoc hasn't installed yet: https://grpc.io/docs/protoc-installation/"))
	}

	protocGenGoVersion, err := exec.Command(`protoc-gen-go`, "--version").Output()
	if err == nil {
		fmt.Printf(SuccessColor, "[v] ")
		fmt.Println(fmt.Sprintf("protoc-gen-go (%s)", fmt.Sprint(strings.Replace(string(protocGenGoVersion), "\n", "", -1))))
	} else {
		fmt.Printf(ErrorColor, "[x] ")
		fmt.Println("protoc-gen-go")
		fmt.Printf(ErrorColor+"\n", fmt.Sprint(strutil.Resize("", 4), "protoc-gen-go hasn't installed yet: https://grpc.io/docs/languages/go/quickstart/"))
	}

	protocGenGoGrpcVersion, err := exec.Command(`protoc-gen-go-grpc`, "--version").Output()
	if err == nil {
		fmt.Printf(SuccessColor, "[v] ")
		fmt.Println(fmt.Sprintf("protoc-gen-go-grpc (%s)", fmt.Sprint(strings.Replace(string(protocGenGoGrpcVersion), "\n", "", -1))))
	} else {
		fmt.Printf(ErrorColor, "[x] ")
		fmt.Println("protoc-gen-go-grpc")
		fmt.Printf(ErrorColor+"\n", fmt.Sprint(strutil.Resize("", 4), "protoc-gen-go-grpc hasn't installed yet: https://grpc.io/docs/languages/go/quickstart/"))
	}

	protocGenGrpcGatewayVersion, err := exec.Command(`protoc-gen-grpc-gateway`, "--version").Output()
	if err == nil {
		fmt.Printf(SuccessColor, "[v] ")
		fmt.Println(fmt.Sprintf("protoc-gen-grpc-gateway (%s)", fmt.Sprint(strings.Replace(string(protocGenGrpcGatewayVersion), "\n", "", -1))))
	} else {
		fmt.Printf(ErrorColor, "[x] ")
		fmt.Println("protoc-gen-grpc-gateway")
		fmt.Printf(ErrorColor+"\n", fmt.Sprint(strutil.Resize("", 4), "protoc-gen-grpc-gateway hasn't installed yet: https://github.com/grpc-ecosystem/grpc-gateway"))
	}

	protocGenCddVersion, err := exec.Command(`protoc-gen-cdd`, "--version").Output()
	if err == nil {
		fmt.Printf(SuccessColor, "[v] ")
		fmt.Println(fmt.Sprintf("protoc-gen-cdd (%s)", fmt.Sprint(strings.Replace(string(protocGenCddVersion), "\n", "", -1))))
	} else {
		fmt.Printf(ErrorColor, "[x] ")
		fmt.Println("protoc-gen-cdd")
		fmt.Printf(ErrorColor+"\n", fmt.Sprint(strutil.Resize("", 4), "protoc-gen-cdd hasn't installed yet: https://github.com/krisnasw/go-grst/tree/main/protoc-gen-cdd"))
	}

	return nil
}
