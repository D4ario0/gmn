package models

import (
	"context"

	"google.golang.org/genai"
)

const (
	GEMINI_2_5_FLASH = "gemini-2.5-flash"
	GEMINI_2_0_FLASH = "gemini-2.0-flash"
)

type GeminiConfig struct {
	model          string
	ctx            context.Context
	thinkingBudget int32
	client         *genai.Client
	systemPrompts  *genai.Content
}

func Init(model, role string, ctx context.Context) (*GeminiConfig, error) {
	geminiClient, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}
	sys_role := genai.NewContentFromText(role, genai.RoleUser)
	return &GeminiConfig{model, ctx, int32(0), geminiClient, sys_role}, nil

}

func PromptModel(prompt string, conf *GeminiConfig) (string, error) {
	response, err := conf.client.Models.GenerateContent(conf.ctx, conf.model, genai.Text(prompt), &genai.GenerateContentConfig{
		ThinkingConfig:    &genai.ThinkingConfig{ThinkingBudget: &conf.thinkingBudget},
		SystemInstruction: conf.systemPrompts,
	})
	return response.Text(), err
}
