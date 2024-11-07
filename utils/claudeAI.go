package utils

import (
	"context"
	"fmt"
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic"
	"github.com/madebywelch/anthropic-go/v3/pkg/anthropic/client/native"
	"log"
	"os"
	"time"
)

type ClaudeAI struct {
	client *native.Client
}

var Claude *ClaudeAI

func NewClaudeAI(apiKey string) (*ClaudeAI, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	client, err := native.MakeClient(native.Config{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Claude client: %w", err)
	}

	return &ClaudeAI{
		client: client,
	}, nil
}

func init() {
	var err error
	if Claude, err = NewClaudeAI(os.Getenv("API_KEY")); err != nil {
		panic("failed to initialize Claude")
	} else {
		log.Println("Claude AI initialized successfully")

	}

}

func (c *ClaudeAI) SendMessage(message string) (string, error) {

	contentBlock := anthropic.NewTextContentBlock(message)

	messagePart := anthropic.MessagePartRequest{
		Role:    "user",
		Content: []anthropic.ContentBlock{contentBlock},
	}

	request := anthropic.NewMessageRequest(
		[]anthropic.MessagePartRequest{messagePart},
		anthropic.WithModel[anthropic.MessageRequest]("claude-3-sonnet-20240229"),
		// 設置合理的最大令牌數
		anthropic.WithMaxTokens[anthropic.MessageRequest](500),
		// 添加系統消息（可選）
		anthropic.WithSystemPrompt[anthropic.MessageRequest]("You are a patent risk assessment expert."),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := c.client.Message(ctx, request)
	if err != nil {
		return "", fmt.Errorf("1failed to get response from Claude: %w", err)
	}

	if len(response.Content) == 0 {
		return "", fmt.Errorf("2received empty response from Claude")
	}

	// 返回響應文本
	return response.Content[0].Text, nil
}
