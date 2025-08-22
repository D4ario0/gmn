package main

import (
	"context"
	"log"
	"os"

	"github.com/D4ario0/gmn/internal/config"
	"github.com/D4ario0/gmn/internal/models"
	"github.com/D4ario0/gmn/internal/tui"
)

const SYS_PROMPT = `- You go straight to the point, brief answers with examples if applicable; unless asked to explain.
- Any question regarding environment configuration assume it is for Linux Fedora 42.
- You answer using markdown format.`

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Print("Could not find a config, gmn will default to `GEMINI 2.0 FLASH`")
		cfg.Model = models.GEMINI_2_0_FLASH
	}

	client, err := models.Init(cfg.Model, cfg.GetProfileConfig(), context.Background())
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

			r, err := models.PromptModel(input, client)
			if err != nil {
				log.Fatal(err)
				return ""
			}
			return r
		}()

		tui.OutputResponse(response, os.Stdout, tui.NewResponseFormatter())

	}

}
