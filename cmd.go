package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

type CMD struct {
	cmd    *exec.Cmd
	stdin  *bytes.Buffer
	stdout *bytes.Buffer
}

func NewCMD(bin string) *CMD {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return nil
	}
	cmd := exec.Command(cwd + "/" + bin)
	stdin := &bytes.Buffer{}
	stdout := &bytes.Buffer{}

	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stdout // Capture standard error as well

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting the command:", err)
		return nil
	}

	return &CMD{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
	}
}

func (c *CMD) Stop() {
	c.cmd.Process.Kill()
}

func (c *CMD) Send(data []byte) []byte {
	_, err := c.stdin.Write(data)
	if err != nil {
		fmt.Println("Error writing to stdin:", err)
	}

	_, err = c.stdin.Write([]byte("\n"))
	if err != nil {
		fmt.Println("Error writing newline to stdin:", err)
	}

	err = c.cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command to finish:", err)
	}

	output := c.stdout.Bytes()
	return output
}

func (c *CMD) Recv() []byte {
	return c.stdout.Bytes()
}
