package chat

import (
	"encoding/json"
	"fmt"

	"github.com/tmc/langchaingo/llms"
)

type History struct {
	Messages []llms.MessageContent `json:"messages"`
}

func (h *History) AddHumanMessage(content string) {
	h.Messages = append(h.Messages, llms.TextParts(
		llms.ChatMessageTypeHuman,
		content,
	))
}

func (h *History) AddAIMessage(content string) {
	h.Messages = append(
		h.Messages,
		llms.TextParts(
			llms.ChatMessageTypeAI,
			content,
		))
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
