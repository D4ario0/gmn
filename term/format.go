package term

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/alecthomas/chroma/quick"
)

const (
	DEFAULT_LEXER     = "markdown"
	DEFAULT_FORMATTER = "terminal16m"
	DEFAULT_STYLE     = "vulcan"
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

func OutputResponse(response string, w io.Writer, fmtConf *FormatConfig) {
	scanner := bufio.NewScanner(formatResponse(response, fmtConf))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			_, _ = fmt.Fprintf(w, "%s ", word)
			time.Sleep(10 * time.Millisecond)
		}
		_, _ = fmt.Fprint(w, "\n") // preserve line break
	}
}
