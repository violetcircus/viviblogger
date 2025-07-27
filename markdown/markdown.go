package markdown

// the main markdown package file

import (
	"bufio"
	"bytes"
	"github.com/violetcircus/viviblogger/output"
	"log"
	"strings"
)

// post title will be first h1. somehow need to figure out what the preview is, too

func Convert(scanner *bufio.Scanner, post *output.Post) string {
	// create builder to read into from the markdown file and parse line-based formatting (lists, headings)
	var builder strings.Builder
	builder.WriteString(`{{ define "content" }}`)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())

		// check for line-specific formatting characters
		if len(line) > 0 {
			switch line[0] {
			case '#':
				handleHeadings(line, scanner.Text(), &builder, post)
				builder.WriteString("\n")
			case ' ':
				if line[1] == '-' {
					// handleList()
				}
			default:
				builder.WriteString(scanner.Text())
			}
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal("scanner error:", err)
	}
	// end the template
	builder.WriteString(`{{ end }}`)

	// start reformatting through replacing in entire string
	buf := handleText(builder.String())

	log.Print(builder.String())
	return buf
}
