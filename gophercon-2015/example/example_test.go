package main

import (
	"runtime"
	"testing"
)

func TestConvertToTheNumber10(t *testing.T) {
	runtime.Breakpoint()
	i := convertToTheNumber10(5)
	if i != 10 {
		t.Fatalf("Expected %d to equal 10", i)
	}
}
