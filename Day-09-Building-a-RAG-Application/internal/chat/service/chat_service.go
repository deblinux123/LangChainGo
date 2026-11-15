package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/chat/domain"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/models"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/monitoring"
	"github.com/deblinux123/LangChainGo/Day-09-Building-a-RAG-Application/internal/subscription"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type ChatService struct {
	llm   *ollama.LLM
	model string

	usageService *monitoring.UsageService

	limitChecker *subscription.LimitChecker
}

func NewChatService(
	model string,
	usageService *monitoring.UsageService,
	limitChecker *subscription.LimitChecker,
) (*ChatService, error) {

	if usageService == nil {
		return nil, fmt.Errorf("❌ usage service cannot be empty.")
	}

	if limitChecker == nil {
		return nil, fmt.Errorf("❌ limit Checker cannot be empty.")
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
		limitChecker: limitChecker,
	}, nil
}

func (c *ChatService) Chat(
	ctx context.Context,
	req domain.ChatRequest,
) (*domain.ChatResponse, error) {
	if strings.TrimSpace(req.Message) == "" {
		return nil, fmt.Errorf("❌ message cannot be empty.")
	}

	if req.UserID == 0 {
		return nil, fmt.Errorf("❌ user ID cannot be zero")
	}

	if err := c.limitChecker.CheckLimits(req.UserID); err != nil {
		return nil, fmt.Errorf("❌ chat request rejected: %w", err)
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
