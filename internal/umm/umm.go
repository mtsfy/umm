package umm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mtsfy/umm/internal/ai"
	"github.com/mtsfy/umm/internal/history"
	"github.com/mtsfy/umm/internal/types"
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

func Execute(id int) {
	var targetInter types.Interaction

	if id == -1 {
		targetInter = history.GetLatest()
	} else {
		targetInter = history.GetByID(id)
	}

	if targetInter.AIResponse.Command == "" {
		fmt.Println("No command suggestion found in history")
		return
	}

	cmdArr := strings.Fields(targetInter.AIResponse.Command)
	if len(cmdArr) == 0 {
		fmt.Println("Invalid command suggestion")
		return
	}

	fmt.Printf("About to run: %s\n", targetInter.AIResponse.Command)
	fmt.Print("Do you want to execute this command? (y/n): ")

	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) != "y" {
		fmt.Println("Command execution canceled.")
		return
	}

	cmd := exec.Command(cmdArr[0], cmdArr[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Running command...")

	if err := cmd.Run(); err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
	}
}
