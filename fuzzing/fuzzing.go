package fuzzing

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Fuzz(data []byte) int {
	_, err := Run(data)
	if err != nil {
		return 0
	}
	return 1

}

func Run(data []byte) (string, error) {
	cmd := exec.Command("../test_bin.exe")
	cmd.Stdin = bytes.NewReader(data)
	output := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = stdErr
	err := cmd.Run()
	fmt.Println(output.String())
	if err != nil {
		return "", fmt.Errorf("error: %s", err)
	}
	if stdErr.Len() > 0 {
		return "", fmt.Errorf("error: %s", stdErr.String())
	}
	fmt.Println(output.String())
	if !strings.Contains(output.String(), "flag") {
		return "", nil
	}
	return fmt.Sprintf("Flag: %s\n", ExtractFlag(output.String())), nil
}
