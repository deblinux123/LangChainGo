package chat

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/langchaingo/llms"
)

type StoredMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type History struct {
	Messages []StoredMessage `json:"messages"`
}

func (h *History) AddHumanMessage(content string) {
	h.Messages = append(h.Messages, StoredMessage{
		Role:    string(llms.ChatMessageTypeHuman),
		Content: content,
	})
}

func (h *History) AddAIMessage(content string) {
	h.Messages = append(
		h.Messages,
		StoredMessage{
			Role:    string(llms.ChatMessageTypeAI),
			Content: content,
		})
}

func (h *History) Marshal() ([]byte, error) {
	data, err := json.Marshal(h)
	if err != nil {
		return nil, fmt.Errorf("marshal history: %w", err)
	}

	return data, nil
}

func (h *History) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, h); err != nil {
		return fmt.Errorf("unmarshal history: %w", err)
	}

	return nil
}

func (h *History) ToLLMMessages() []llms.MessageContent {
	messages := make([]llms.MessageContent, 0, len(h.Messages))

	for _, msg := range h.Messages {
		messages = append(messages, llms.TextParts(
			llms.ChatMessageType(msg.Role),
			msg.Content,
		))
	}

	return messages
}

func (h *History) AddSystemMessage(content string) {
	h.Messages = append(h.Messages, StoredMessage{
		Role:    string(llms.ChatMessageTypeSystem),
		Content: content,
	})
}
