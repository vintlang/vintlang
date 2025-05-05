package utils

import (
	"os"
	"os/exec"
)

func RunShell(cmd string) error {
	shell := []string{"sh", "-c", cmd}
	if os.Getenv("OS") == "Windows_NT" {
		shell = []string{"cmd", "/C", cmd}
	}
	c := exec.Command(shell[0], shell[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
