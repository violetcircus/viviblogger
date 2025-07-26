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
	// handle html formatting
	if len(line) > 0 {
		handleHeadings(line, content, builder, post)
	}
	// add line breaks after first line
	if builder.String() != `{{ define "content" }}` {
		builder.WriteString("<br>")
	}
	builder.WriteString("\n")
	return builder.String()
}
