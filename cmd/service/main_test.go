package main

import (
	"testing"
)

func TestInc(t *testing.T) {
	if Inc(2, 2) != 4 {
		t.Error("Expected 2+2 to equel 4")
	}
}
