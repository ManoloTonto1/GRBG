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

	// Parse flags
	flag.Parse()

	if *bin == "" {
		fmt.Println("Please provide a binary path")
		os.Exit(1)
	}

	globalScope.binaryPath = *bin
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
