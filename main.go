package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
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
		fmt.Printf("Kamal Proxy Log: %s\n", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading logs: %v\n", err)
	}
}
