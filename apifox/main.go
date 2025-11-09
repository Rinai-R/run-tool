package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("/opt/apifox-2.7.48/apifox",
		"--ozone-platform-hint=auto",
		"--enable-wayland-ime",
		"--wayland-text-input-version=3",
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
