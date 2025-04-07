package main

import (
	"fmt"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {
	newFilePath := filePath + ".processing"

	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "+faststart", "-f", "mp4", newFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("FFmpeg error: %v\nOutput: %s\n", err, string(output))
		return "", err
	}

	return newFilePath, nil
}
