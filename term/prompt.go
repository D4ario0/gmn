package term

import (
	"bufio"
	"fmt"
	"io"
)

const USER_PROMPT = "\033[38;2;80;200;120mgemini ❱❱ \033[0m"

func PromptUser(scanner *bufio.Scanner, w io.Writer) (string, bool) {
	_, _ = fmt.Fprint(w, USER_PROMPT)

	if scanner.Scan() {
		return scanner.Text(), true
	}
	return "", false
}
