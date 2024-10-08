package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	endpoint := os.Getenv("SIDECAR_ENDPOINT")
	if endpoint == "" {
		fmt.Println("Error: SIDECAR_ENDPOINT environment variable not set")
		return
	}

	cmd := exec.Command("docker", "logs", "-f", "kamal-proxy")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error accessing logs: %v\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting docker logs command: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer([]byte(line)))
		if err != nil {
			fmt.Printf("Error sending log to endpoint: %v\n", err)
			continue
		}
		resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading logs: %v\n", err)
	}
}
