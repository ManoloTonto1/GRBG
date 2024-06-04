package main

import (
	"os"
	"os/exec"
	"testing"
)

func run() error {
	program := "temp/fuzzer_test.exe"
	// Build the program
	cmd := exec.Command("go", "build", "-o", program, ".")
	err := cmd.Run()
	if err != nil {
		return err
	}
	// Run the program
	cmd = exec.Command(program, "-bin", "./test_bin.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func TestBuildAndRunProgram(t *testing.T) {
	err := run()
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkProgram(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := run()
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()

}
