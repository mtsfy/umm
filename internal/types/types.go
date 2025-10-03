package types

import "time"

type AIResponse struct {
	Description string `json:"description"`
	Command     string `json:"command"`
}

type Interaction struct {
	Model        string        `json:"model"`
	ResponseTime time.Duration `json:"response_time"`
	TokensUsed   int           `json:"tokens_used"`
	Date         time.Time     `json:"date"`
	UserInput    string        `json:"user_input"`
	AIResponse   AIResponse    `json:"ai_response"`
}

type History struct {
	Interactions []Interaction `json:"interactions"`
}

type Config struct {
	ApiKey string `json:"api_key"`
	Model  string `json:"model"`
}
