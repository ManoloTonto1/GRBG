package fuzzing_test

import (
	"testing"

	"github.com/manolotonto1/GRBG/fuzzing"
)

func TestRun(t *testing.T) {
	data := []byte("testing 1,2,3")
	flag, err := fuzzing.Run(data)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if flag != "" {
		t.Errorf("Flag: %s", flag)
	}

}
