package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	SUCCESS_STRING = iota
	FAIL_STRING
)

type GlobalScope struct {
	flags      map[int]string
	binaryPath string
}

var globalScope = GlobalScope{
	flags:      make(map[int]string),
	binaryPath: "",
}

func main() {

	// Define flags
	bin := flag.String(
		"bin",
		"",
		"Path to the binary to fuzz",
	)

	success := flag.String(
		"success",
		"",
		"String to search for in the response body, this indicates a successful fuzzing attempt",
	)
	fail := flag.String(
		"fail",
		"",
		"String to search for in the response body, this indicates a failed fuzzing attempt",
	)

	// Parse flags
	flag.Parse()

	if *success == "" {
		fmt.Println("Please provide a success string")
		os.Exit(1)
	}
	if *fail == "" {
		fmt.Println("Please provide a fail string")
		os.Exit(1)
	}
	if *bin == "" {
		fmt.Println("Please provide a binary path")
		os.Exit(1)
	}

	globalScope.flags[SUCCESS_STRING] = *success
	globalScope.flags[FAIL_STRING] = *fail
	globalScope.binaryPath = *bin
	println("Success string: ", globalScope.flags[SUCCESS_STRING])
	println("Fail string: ", globalScope.flags[FAIL_STRING])
	// Start the fuzzer
	exploiters := []Exploiter{
		NewFormatStringExploiter(),
	}
	for _, exploiter := range exploiters {
		if success := exploiter.Exploit(); success {
			break
		}
	}
}
