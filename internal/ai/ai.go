package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/mtsfy/umm/internal/history"
	"github.com/mtsfy/umm/internal/types"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client openai.Client
var system string

func Ask(query string) {
	prompt := fmt.Sprintf(`You are a technical CLI assistant. Respond with a technical, short answer in JSON format containing two fields: 'description' for a brief summary and 'command' for an example command to run. Ensure that the output is strictly valid JSON. The system is %s.`, system)

	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(prompt),
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
	system = runtime.GOOS

	apiKey := os.Getenv("UMM_AI_TOKEN")

	if apiKey == "" {
		fmt.Println("Unable to get OpenAI API Key.")
		os.Exit(1)
	}

	client = openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}
