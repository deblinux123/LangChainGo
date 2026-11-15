package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/domain"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type ChatService struct {
	llm   *ollama.LLM
	model string
}

func NewModelService(model string) (*ChatService, error) {
	llm, err := ollama.New(
		ollama.WithModel(model),
	)

	if err != nil {
		return nil, fmt.Errorf("❌ initialize ollama: %w", err)
	}

	return &ChatService{
		llm:   llm,
		model: model,
	}, nil
}

func (c *ChatService) Chat(
	ctx context.Context,
	req domain.ChatRequest,
) (*domain.ChatResponse, error) {
	if strings.TrimSpace(req.Message) == "" {
		return nil, fmt.Errorf("❌ message cannot be empty.")
	}

	messages := []llms.MessageContent{
		llms.TextParts(
			llms.ChatMessageTypeHuman,
			req.Message,
		),
	}

	response, err := c.llm.GenerateContent(ctx, messages)
	fmt.Printf("Raw Ollama Response: %#v\n", response)

	if err != nil {
		return nil, fmt.Errorf("❌ generate chat response: %w", err)
	}

	if response == nil || len(response.Choices) == 0 {
		return nil, fmt.Errorf("🔴 ollama returned an empty response")
	}

	answer := response.Choices[0]

	generateInfo := answer.GenerationInfo
	promptToken, _ := generateInfo["PromptTokens"].(int)
	CompleTokens, _ := generateInfo["CompletionTokens"].(int)
	totalToken, _ := generateInfo["TotalTokens"].(int)

	return &domain.ChatResponse{
		Answer:           answer.Content,
		Model:            c.model,
		PromptTokens:     int64(promptToken),
		CompletionTokens: int64(CompleTokens),
		TotalTokens:      int64(totalToken),
	}, nil
}
