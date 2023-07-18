package gocli

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetCurrentModule() (string, error) {
	cmd := exec.Command(`go`, "list", "-m")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}
	return strings.ReplaceAll(out.String(), "\n", ""), nil
}
