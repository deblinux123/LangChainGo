package chat

import (
	"context"
	"fmt"

	"github.com/deblinux123/LangChainGo/Day-05-ConversationMemory/chatbot/internal/redisclient"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Service struct {
	model llms.Model
	redis *redisclient.Client
}

func NewService(model llms.Model, redis *redisclient.Client) *Service {
	return &Service{
		model: model,
		redis: redis,
	}
}

func (s *Service) Chat(ctx context.Context, sessionID string, userMessage string) (string, error) {
	if sessionID == "" {
		return "", fmt.Errorf("SessionID cannot be empty")
	}

	if userMessage == "" {
		return "", fmt.Errorf("Message cannot be empty.")
	}

	history := &History{}

	// load previous history

	data, err := s.redis.GetHistory(ctx, sessionID)

	if err != nil {
		return "", err
	}

	if len(data) > 0 {
		if err := history.Unmarshal(data); err != nil {
			return "", err
		}
	} else {
		history.Messages = []llms.MessageContent{
			llms.TextParts(
				llms.ChatMessageTypeSystem,
				"You are a helpful assistant. Answer clearly in persian and remember the conversation context.",
			),
		}
	}

	history.AddHumanMessage(userMessage)

	response, err := s.model.GenerateContent(ctx, history.Messages)

	if err != nil {
		return "", fmt.Errorf("generate response: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("model return no response.")
	}

	answer := response.Choices[0].Content

	history.AddAIMessage(answer)

	data, err = history.Marshal()

	if err != nil {
		return "", err
	}

	if err := s.redis.SaveHistory(ctx, sessionID, data); err != nil {
		return "", err
	}

	return answer, nil
}

var _ = ollama.New
