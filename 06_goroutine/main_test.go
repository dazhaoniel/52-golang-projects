package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printNames(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printNames("Test", &wg)

	wg.Wait()

	_ = w.Close()

	res, _ := io.ReadAll(r)

	output := string(res)

	os.Stdout = stdOut

	if !strings.Contains(output, "Test") {
		t.Errorf("Expected Test but got %s", output)
	}
}
