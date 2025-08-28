package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mtsfy/umm/internal/config"
	"github.com/mtsfy/umm/internal/history"
	"github.com/mtsfy/umm/internal/types"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client openai.Client

func Ask(query string) {
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage("You are a technical CLI assistant. Respond with a technical, short answer in JSON format containing two fields: 'description' for a brief summary and 'command' for an example command to run. Ensure that the output is strictly valid JSON."),
				openai.UserMessage(query),
			},
			Model: openai.ChatModelGPT4oMini,
		},
	)
	if err != nil {
		panic(err)
	}

	content := chatCompletion.Choices[0].Message.Content
	res := parseResponse(content)

	fmt.Println(res.Description)
	fmt.Println(res.Command)

	history.Save(types.Interaction{
		UserInput:  query,
		AIResponse: res,
	})
}

func parseResponse(content string) types.AIResponse {
	var res types.AIResponse

	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimSuffix(content, "```")

	json.Unmarshal([]byte(content), &res)
	return res
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
