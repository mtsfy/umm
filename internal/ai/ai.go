package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mtsfy/umm/internal/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type Response struct {
	Description string `json:"description"`
	Command     string `json:"command"`
}

var client openai.Client

func Ask(query string) {
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage("You are a technical CLI assistant. Respond with a technical, short answer in JSON format containing two fields: 'description' for a brief summary and 'command' for an example command to run. Ensure that the output is strictly valid JSON."),
				openai.UserMessage(query),
			},
			Model: openai.ChatModelGPT4o,
		},
	)
	if err != nil {
		panic(err)
	}

	content := chatCompletion.Choices[0].Message.Content
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimSuffix(content, "```")

	var r Response
	json.Unmarshal([]byte(content), &r)
	fmt.Println(r.Description)
	fmt.Println(r.Command)
}

func init() {
	apiKey := config.Config("OPENAI_API_KEY")

	if apiKey == "" {
		fmt.Println("Unable to get OpenAI API Key.")
		os.Exit(1)
	}

	client = openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}
