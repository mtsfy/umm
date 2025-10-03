package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/mtsfy/umm/internal/config"
	"github.com/mtsfy/umm/internal/history"
	"github.com/mtsfy/umm/internal/types"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client openai.Client
var system string

func Ask(query string) error {
	start := time.Now()
	prompt := fmt.Sprintf(`You are a technical CLI assistant. Respond with a technical, short answer in JSON format containing two fields: 'description' for a brief summary and 'command' for an example command to run. Ensure that the output is strictly valid JSON. The system is %s.`, system)

	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(prompt),
				openai.UserMessage(query),
			},
			Model: getModel(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create chat completion: %w", err)
	}

	content := chatCompletion.Choices[0].Message.Content

	res, err := parseResponse(content)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Println(res.Description)
	fmt.Println(res.Command)

	responseTime := time.Since(start)
	modelName := config.Get("MODEL")

	if err := history.Save(types.Interaction{
		Date:         start,
		UserInput:    query,
		ResponseTime: responseTime,
		Model:        modelName,
		AIResponse:   res,
		TokensUsed:   int(chatCompletion.Usage.TotalTokens),
	}); err != nil {
		return fmt.Errorf("failed to save interaction to history: %w", err)
	}

	return nil
}

func FollowUp(lastInteraction types.Interaction, query string) error {
	date := time.Now()
	prompt := fmt.Sprintf(`You are a technical CLI assistant. Respond with a technical, concise answer in JSON format containing two fields: "description" for a brief summary and "command" for an example command to run. Ensure that the output is strictly valid JSON. 
The previous query was: "%s"
The suggested command was: "%s"
The system is %s.`,
		lastInteraction.UserInput,
		lastInteraction.AIResponse.Command,
		system,
	)

	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(prompt),
				openai.UserMessage(query),
			},
			Model: getModel(),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create chat completion: %w", err)
	}

	content := chatCompletion.Choices[0].Message.Content

	res, err := parseResponse(content)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	fmt.Println(res.Description)
	fmt.Println(res.Command)

	if err := history.Save(types.Interaction{
		Date:       date,
		UserInput:  query,
		AIResponse: res,
	}); err != nil {
		return fmt.Errorf("failed to save interaction to history: %w", err)
	}

	return nil
}

func parseResponse(content string) (types.AIResponse, error) {
	var res types.AIResponse

	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	if err := json.Unmarshal([]byte(content), &res); err != nil {
		return res, err
	}
	return res, nil
}

func getModel() openai.ChatModel {
	modelName := config.Get("MODEL")
	switch modelName {
	case "gpt-4o-mini":
		return openai.ChatModelGPT4oMini
	case "gpt-4o":
		return openai.ChatModelGPT4o
	case "gpt-4":
		return openai.ChatModelGPT4
	default:
		return openai.ChatModelGPT4oMini
	}
}

func init() {
	system = runtime.GOOS

	apiKey := config.Get("API_KEY")

	if apiKey == "" {
		fmt.Println("Unable to get OpenAI API Key.")
		os.Exit(1)
	}

	client = openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}
