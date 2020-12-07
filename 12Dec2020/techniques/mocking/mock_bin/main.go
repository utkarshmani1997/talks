package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

var (
	execCommand = exec.Command
)

func main() {
	files, err := listAllFiles()
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}

func listAllFiles() (string, error) {
	var stdout bytes.Buffer
	cmd := execCommand("ls", "-l")
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return string(stdout.Bytes()), nil
}
