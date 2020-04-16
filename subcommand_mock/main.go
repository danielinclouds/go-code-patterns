package main

import (
	"os/exec"
	"strings"
)

var execCommand = exec.Command

// KindGetClusters func
func KindGetClusters() string {
	b, err := execCommand("kind", "get", "clusters").CombinedOutput()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(b))
}

func main() {
	_ = KindGetClusters()
}
