package umm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mtsfy/umm/internal/ai"
	"github.com/mtsfy/umm/internal/history"
)

func Query(q string) {
	if q == "" {
		fmt.Println("empty query")
		return
	}

	ai.Ask(q)
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
