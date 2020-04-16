package main

import (
	"os/exec"
	"testing"
)

func TestKindGetClusters_mock(t *testing.T) {

	execCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "No kind clusters found.")
	}
	defer func() {
		execCommand = exec.Command
	}()

	got := KindGetClusters()
	want := "No kind clusters found."
	if got != want {
		t.Errorf("Echo() = %q; want %q", got, want)
	}
}
