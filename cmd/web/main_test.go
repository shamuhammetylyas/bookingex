package main

import "testing"

// main package-in icindaki run functionyny test etmek ucin yazyldy.
// ady hokman TestRun bolmak hokman dal yone asy hokman Test sozi bilen bashlamalydyr.
func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}
