package umm

import (
	"fmt"

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

	fmt.Println("running last suggested command")
	fmt.Printf("command: %s\n", latest.AIResponse.Command)
}
