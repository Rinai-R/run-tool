package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("Usage: run-code <path>")
	}

	cmd := exec.Command("/usr/bin/code",
		"--ozone-platform-hint=auto",
		"--enable-wayland-ime",
		"--wayland-text-input-version=3",
		args[0],
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
