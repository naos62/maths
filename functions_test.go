package main

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	if Pow(2) != 4 {
		fmt.Println("Expected 4")
		t.Fail()
	}
	if Pow(3) != 8 {
		fmt.Println("Expected 9")
		t.Fail()
	}
}
