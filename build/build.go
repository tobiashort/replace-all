package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "build", "-o", "build/replace-all")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
