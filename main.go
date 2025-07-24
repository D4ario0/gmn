package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"piplup/models"
	"piplup/term"
)

const SYS_PROMPT = `
	- You go straight to the point and give brief answers, oneliners followed by an example if applicable; unless asked to explain.
	- Any question regarding environment configuration assume it is for Linux Fedora 42.
	- You answer using markdown format.
`

func main() {
	config, err := models.Init(models.GEMINI_2_5_FLASH, SYS_PROMPT, context.Background())
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		input, ok := term.PromptUser(scanner, os.Stdout)
		if !ok || input == "exit" {
			break
		}

		response, err := models.PromptModel(input, config)
		if err != nil {
			log.Fatal(err)
		}

		term.OutputResponse(response, os.Stdout, nil)

	}

}
