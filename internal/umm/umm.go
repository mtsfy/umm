package umm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mtsfy/umm/internal/ai"
	"github.com/mtsfy/umm/internal/history"
)

func Query(query string) {
	query = strings.TrimSpace(query)
	if query == "" {
		fmt.Println("No query provided. Please enter a valid query.")
		return
	}

	if err := ai.Ask(query); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing query: %v\n", err)
	}
}

func Execute() {
	latest := history.GetLatest()
	if latest.AIResponse.Command == "" {
		fmt.Println("No command suggestion found in history")
		return
	}

	cmdArr := strings.Fields(latest.AIResponse.Command)
	if len(cmdArr) == 0 {
		fmt.Println("Invalid command suggestion")
		return
	}

	cmd := exec.Command(cmdArr[0], cmdArr[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Running last suggested command")
	fmt.Printf("Command: %s\n", latest.AIResponse.Command)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
	}
}
