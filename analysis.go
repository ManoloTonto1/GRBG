package main

const (
	READ = iota
	WRITE
	ERROR
	EXIT
)

type IOCycle []int

// AnalyzeIO runs the binary and analyzes the I/O cycles
// find out what the flow of input and output is like
// and return a slice of the ENUM representing the I/O cycles
func AnalyzeIO() IOCycle {
	cycle := IOCycle{}
	cmd := NewCMD(globalScope.binaryPath)
	for {
		// try to read from the binary
		if output := cmd.Recv(); output != nil {
			cycle = append(cycle, READ)
			continue
		}
		// Get the output from the binary
		if output := cmd.Send([]byte("A")); output != nil {
			cycle = append(cycle, WRITE)
			continue
		}
		// Check if the binary errored out
		if output := cmd.Send([]byte("A")); output != nil {
			cycle = append(cycle, ERROR)
			continue
		}
		// Check if the binary exited
		if output := cmd.Send([]byte("A")); output != nil {
			cycle = append(cycle, EXIT)
			break
		}

	}
	return cycle
}
