package main

import (
	"context"
	"log"
	"os"

	"github.com/D4ario0/gmn/internal/config"
	"github.com/D4ario0/gmn/internal/models"
	"github.com/D4ario0/gmn/internal/tui"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Print("Could not find a config, gmn will default to `GEMINI 2.0 FLASH`")
	}

	chatInstance, err := models.GetChatter(cfg.Model, cfg.GetProfileConfig(), context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for {
		input, ok := tui.PromptUser(os.Stdin, os.Stdout)
		if !ok || input == "" || input == "exit" {
			break
		}

		response := func() string {
			sp := tui.NewSpinner("Generating LLM response...").Start()
			defer sp.Stop()

			r, err := chatInstance.PromptModel(input)
			if err != nil {
				log.Fatal(err)
				return ""
			}
			return r
		}()

		tui.OutputResponse(response, os.Stdout, tui.NewResponseFormatter())

	}

}
