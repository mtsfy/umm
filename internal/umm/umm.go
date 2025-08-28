package umm

import (
	"fmt"

	"github.com/mtsfy/umm/internal/ai"
)

func Query(q string) {
	if q == "" {
		fmt.Println("empty query")
		return
	}

	ai.Ask(q)
}

func Execute() {
	fmt.Println("running last suggested command")
}
