package main

import "testing"

// A meaningless test added for fun
func TestMain(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("got = %d; want 1", got)
	}
}
