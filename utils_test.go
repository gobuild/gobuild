package main

import "testing"

func Test(t *testing.T) {
	s := RandNString(10)
	if len(s) != 10 {
		t.Fatalf("expect rand string length 10, but got %d", len(s))
	}
}
