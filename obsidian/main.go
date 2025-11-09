package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/var/lib/snapd/snap/obsidian/50/obsidian",
		"--ozone-platform-hint=auto",
		"--enable-wayland-ime",
		"--wayland-text-input-version=3",
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
