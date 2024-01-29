package main

import (
	"os"
	"io"
	"testing"
	"strings"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout 
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "final bank balance: $34320.00") {
		t.Errorf("expected $34320.00 in output, got %s", output)
	}
}