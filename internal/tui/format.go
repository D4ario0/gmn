package tui

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/alperdrsnn/clime"
)

const (
	USER_PROMPT          = "\033[38;2;80;200;120m GEMINI ❱❱ \033[0m"
	DEFAULT_LEXER        = "markdown"
	DEFAULT_FORMATTER    = "terminal16m"
	DEFAULT_STYLE        = "evergarden"
	OUTPUT_TIME_INTERVAL = 10 * time.Millisecond
)

type ResponseFormatter struct {
	Lexer       string
	Formatter   string
	Style       string
	OutputDelay time.Duration
}

func NewResponseFormatter() *ResponseFormatter {
	return &ResponseFormatter{
		Lexer:       DEFAULT_LEXER,
		Formatter:   DEFAULT_FORMATTER,
		Style:       DEFAULT_STYLE,
		OutputDelay: OUTPUT_TIME_INTERVAL,
	}
}

func (rf *ResponseFormatter) format(response string) *bytes.Buffer {
	responseBuffer := &bytes.Buffer{}
	_ = quick.Highlight(responseBuffer, response, rf.Lexer, rf.Formatter, rf.Style)
	return responseBuffer
}

func PromptUser(r io.Reader, w io.Writer) (string, bool) {
	fmt.Fprint(w, USER_PROMPT)

	// Read all bytes directly from the provided reader until EOF.
	inputBytes, err := io.ReadAll(r)
	if err != nil {
		fmt.Fprintf(w, "Error reading input: %v\n", err)
		return "", false
	}

	return strings.TrimSpace(string(inputBytes)), true
}

func OutputResponse(response string, w io.Writer, fmtConf *ResponseFormatter) {
	if response == "" {
		return
	}

	insertLineBreak := func() { fmt.Fprint(w, "\n") }

	// Print line breaks at the start and at the end
	insertLineBreak()
	defer insertLineBreak()

	formattedOutput := fmtConf.format(response)

	scanner := bufio.NewScanner(formattedOutput)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// Use it to respect tabs as opposed to strings.Fields()
		line := strings.Split(scanner.Text(), " ")
		typeWriter(w, line, fmtConf.OutputDelay)
		insertLineBreak()
	}
}

func typeWriter(w io.Writer, line []string, t time.Duration) {
	for _, word := range line {
		_, _ = fmt.Fprintf(w, "%s ", word)
		time.Sleep(t)
	}
}

func NewSpinner(msg string) *clime.Spinner {
	return clime.NewSpinner().
		WithColor(clime.BlueColor).
		WithStyle(clime.SpinnerDots).
		WithMessage(msg)
}
