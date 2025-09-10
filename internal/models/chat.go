package models

import (
	"context"
	"errors"
)

type Chatter interface {
	PromptModel(p string) (r string, err error)
}

// Supported Models
const (
	GEMINI_2_0_FLASH      = "gemini-2.0-flash"
	GEMINI_2_5_FLASH      = "gemini-2.5-flash"
	GEMINI_2_5_FLASH_LITE = "gemini-2.5-flash-lite"
)

func GetChatter(model, sysPrompt string, ctx context.Context) (Chatter, error) {
	switch model {
	case GEMINI_2_0_FLASH, GEMINI_2_5_FLASH, GEMINI_2_5_FLASH_LITE:
		return newGeminiChat(model, sysPrompt, ctx)
	// Add more models here
	default:
		return nil, errors.New("unsupported model")
	}
}
