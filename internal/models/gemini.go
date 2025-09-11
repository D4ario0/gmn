package models

import (
	"context"

	"google.golang.org/genai"
)

type GeminiChat struct {
	ctx        context.Context
	model      string
	chatClient *genai.Chat
}

func newGeminiChat(model, systemPrompt string, ctx context.Context) (Chatter, error) {
	var history []*genai.Content

	geminiClient, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}

	chat, err := geminiClient.Chats.Create(ctx, model, &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
		ThinkingConfig:    &genai.ThinkingConfig{ThinkingBudget: nil},
	}, history)

	if err != nil {
		return nil, err
	}

	return &GeminiChat{ctx, model, chat}, nil
}

func (c *GeminiChat) PromptModel(prompt string) (string, error) {
	if !checkInternetConnection() {
		return "**No internet connection. Please make sure you are connected to a reliable network**", nil
	}

	response, err := c.chatClient.SendMessage(c.ctx, genai.Part{Text: prompt})

	return response.Candidates[0].Content.Parts[0].Text, err
}
