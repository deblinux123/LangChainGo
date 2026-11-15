package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/domain"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/monitoring"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type ChatService struct {
	llm   *ollama.LLM
	model string

	usageService *monitoring.UsageService
}

func NewChatService(model string, usageService *monitoring.UsageService) (*ChatService, error) {

	if usageService == nil {
		return nil, fmt.Errorf("❌ usage service cannot be empty.")
	}

	llm, err := ollama.New(
		ollama.WithModel(model),
	)

	if err != nil {
		return nil, fmt.Errorf("❌ initialize ollama: %w", err)
	}

	return &ChatService{
		llm:          llm,
		model:        model,
		usageService: usageService,
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

	startTime := time.Now()

	response, err := c.llm.GenerateContent(ctx, messages)

	if err != nil {
		return nil, fmt.Errorf("❌ generate chat response: %w", err)
	}

	if response == nil || len(response.Choices) == 0 {
		return nil, fmt.Errorf("🔴 ollama returned an empty response")
	}

	answer := response.Choices[0]

	if answer == nil {
		return nil, fmt.Errorf("🔴 ollama returned a nil choice")
	}

	generateInfo := answer.GenerationInfo
	promptToken, _ := generateInfo["PromptTokens"].(int)
	CompleTokens, _ := generateInfo["CompletionTokens"].(int)
	totalToken, _ := generateInfo["TotalTokens"].(int)

	usage := &models.Usage{
		UserID:       req.UserID,
		Feature:      "chat",
		Model:        c.model,
		InputTokens:  int64(promptToken),
		OutputTokens: int64(CompleTokens),
		LatencyMs:    time.Since(startTime).Milliseconds(),
		Status:       models.UsageSuccess,
	}

	if err := c.usageService.RecordUsage(usage); err != nil {
		return nil, fmt.Errorf("🔴 record chat usage: %w", err)
	}

	return &domain.ChatResponse{
		Answer:           answer.Content,
		Model:            c.model,
		PromptTokens:     int64(promptToken),
		CompletionTokens: int64(CompleTokens),
		TotalTokens:      int64(totalToken),
	}, nil
}
