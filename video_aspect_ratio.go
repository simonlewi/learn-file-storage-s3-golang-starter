package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func getVideoAspectRatio(filePath string) (string, error) {
	type Stream struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}

	type Response struct {
		Streams []Stream `json:"streams"`
	}

	// Use ffprobe to get the video aspect ratio
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)

	var outb bytes.Buffer
	cmd.Stdout = &outb

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	var resp Response
	if err := json.Unmarshal(outb.Bytes(), &resp); err != nil {
		return "", err
	}

	if len(resp.Streams) == 0 {
		return "", fmt.Errorf("no streams found in video")
	}

	width := resp.Streams[0].Width
	height := resp.Streams[0].Height

	// Calculate ratio as a float
	ratio := float64(width) / float64(height)

	// Compare with tolerance
	if ratio >= 1.7 && ratio <= 1.8 { // 16:9 is approximately 1.78
		return "16:9", nil
	} else if ratio >= 0.55 && ratio <= 0.57 { // 9:16 is approximately 0.56
		return "9:16", nil
	} else {
		return "other", nil
	}
}
