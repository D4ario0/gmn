package tui

import (
	"fmt"
	"io"
)

const USER_PROMPT = "\033[38;2;80;200;120mgemini ❱❱ \033[0m"

func PromptUser(r io.Reader, w io.Writer) (string, bool) {
	fmt.Fprint(w, USER_PROMPT)

	// Read all bytes directly from the provided reader until EOF.
	inputBytes, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(w, "Error reading input: %v\n", err)
		return "", false
	}

	return string(inputBytes), true
}
