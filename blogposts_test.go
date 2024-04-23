package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestPostGenerator(t *testing.T) {
	// get current dir
	cwd, _ := os.Getwd()
	// remove dest dir content
	exec.Command(fmt.Sprintf("rm %s/%s", cwd, "dest/*"))
	// generate templates to files in dest
	PostGenerator()

	// files names based on my knowldge
	paths := []string{"dest/index.html", "dest/welcome-to-my-blog.html"}

	// test files exist or fail test if not exist
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Fatalf("file not exist %s", path)
		}
	}
}
