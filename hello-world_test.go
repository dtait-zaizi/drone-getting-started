package main

import "testing"

func TestSayHi(t *testing.T) {
	result := sayhi()
	if result != "hi" {
		t.Errorf("sayhi() was incorrect, got: %s, want: %s.", result, "hi")
	}
}
