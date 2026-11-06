package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type ChatBot struct {
	model   llms.Model
	history []llms.MessageContent
}

func NewChatBot(model llms.Model) *ChatBot {
	return &ChatBot{
		model: model,
		history: []llms.MessageContent{
			llms.TextParts(
				llms.ChatMessageTypeSystem,
				"you are helpful assistant. Remember the conversation context. and answer in persian language.",
			),
		},
	}
}

func (c *ChatBot) Chat(ctx context.Context, userInput string) (string, error) {
	c.history = append(c.history, llms.TextParts(
		llms.ChatMessageTypeHuman,
		userInput,
	))

	c.history = trimHistory(c.history, 4)
	response, err := c.model.GenerateContent(
		ctx,
		c.history,
	)

	if err != nil {
		c.history = c.history[:len(c.history)-1]

		return "", fmt.Errorf("Generate response failed: %v", err)
	}

	if len(response.Choices) == 0 {
		c.history = c.history[:len(c.history)-1]

		return "", fmt.Errorf("No response generated")
	}

	answer := response.Choices[0].Content

	c.history = append(c.history, llms.TextParts(
		llms.ChatMessageTypeAI,
		answer,
	))

	return answer, nil
}

func main() {

	ctx := context.Background()

	model, err := ollama.New(
		ollama.WithModel("gemma4:e2b"),
	)

	if err != nil {
		log.Fatal(err)
	}

	chatbot := NewChatBot(model)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the chatbot! Type 'exit' to quit.")

	for {
		fmt.Println("\nYou:")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			break
		}

		if input == "" {
			continue
		}

		answer, err := chatbot.Chat(
			ctx,
			input,
		)

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Assistant:", answer)
		fmt.Println("History messages:", len(chatbot.history))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Input Error: %v", err)
	}

}

func trimHistory(history []llms.MessageContent, windowSize int) []llms.MessageContent {
	if len(history) <= 1 {
		return history
	}

	systemMessage := history[0]
	conversationMessage := history[1:]

	maxMessages := windowSize * 4

	if len(conversationMessage) > maxMessages {
		conversationMessage = conversationMessage[len(conversationMessage)-maxMessages:]
	}

	result := make([]llms.MessageContent, 0, len(conversationMessage)+1)
	result = append(result, systemMessage)
	result = append(result, conversationMessage...)

	return result
}
