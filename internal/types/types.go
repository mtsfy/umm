package types

type AIResponse struct {
	Description string `json:"description"`
	Command     string `json:"command"`
}

type Interaction struct {
	UserInput  string     `json:"user_input"`
	AIResponse AIResponse `json:"ai_response"`
}

type History struct {
	Interactions []Interaction `json:"interactions"`
}
