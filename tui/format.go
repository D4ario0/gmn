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
	DEFAULT_LEXER     = "markdown"
	DEFAULT_FORMATTER = "terminal16m"
	DEFAULT_STYLE     = "evergarden"
)

type FormatConfig struct {
	Lexer     string
	Formatter string
	Style     string
}

func formatResponse(response string, fmtConf *FormatConfig) *bytes.Buffer {
	if fmtConf == nil {
		fmtConf = &FormatConfig{
			Lexer:     DEFAULT_LEXER,
			Formatter: DEFAULT_FORMATTER,
			Style:     DEFAULT_STYLE,
		}
	}

	responseBuffer := &bytes.Buffer{}
	_ = quick.Highlight(responseBuffer, response, fmtConf.Lexer, fmtConf.Formatter, fmtConf.Style)
	return responseBuffer
}

// Use it to respect tabs as opposed to strings.Fields()
func splitRespectingTabs(s string) []string {
	return strings.Split(s, " ")
}

func OutputResponse(response string, w io.Writer, fmtConf *FormatConfig) {
	if response == "" {
		return
	}

	scanner := bufio.NewScanner(formatResponse(response, fmtConf))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		words := splitRespectingTabs(line)

		for _, word := range words {
			_, _ = fmt.Fprintf(w, "%s ", word)
			time.Sleep(10 * time.Millisecond)
		}
		_, _ = fmt.Fprint(w, "\n") // preserve line break
	}
}

func NewSpinner(msg string) *clime.Spinner {
	return clime.NewSpinner().
		WithColor(clime.BlueColor).
		WithStyle(clime.SpinnerDots).
		WithMessage(msg)
}
