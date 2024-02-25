package controller

import (
	// "fmt"
	// "os"
	"os/exec"
)

func StartRecord() {
	args := []string{"-r", "48000", "-b", "32", "data/rec.wav", "silence", "-l", "1", "3.0", "0.01%", "2", "50", "0.25%"}
	cmd := exec.Command("rec", args...)
	cmd.Run()
}
