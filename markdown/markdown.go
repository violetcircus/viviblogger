package markdown

// the main markdown package file

import (
	"bufio"
	"bytes"
	"github.com/violetcircus/viviblogger/output"
	"strings"
)

// post title will be first h1. somehow need to figure out what the preview is, too

func Convert(content string, scanner *bufio.Scanner, builder *strings.Builder, post *output.Post) string {
	line := bytes.TrimSpace(scanner.Bytes())
	if len(line) > 0 {
		handleHeadings(line, content, builder, post)
	}
	return builder.String()
}
