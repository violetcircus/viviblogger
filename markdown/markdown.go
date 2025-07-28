package markdown

// the main markdown package file

import (
	"bufio"
	"bytes"
	"github.com/violetcircus/viviblogger/output"
	"log"
	"os"
	"regexp"
	"strings"
)

// post title will be first h1. somehow need to figure out what the preview is, too

func Read(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open %s, %v", fileName, err)
	}
	defer file.Close()

	var post output.Post
	scanner := bufio.NewScanner(file)
	post.Body = convert(scanner, &post)

	output.Build(post)
}

func convert(scanner *bufio.Scanner, post *output.Post) string {
	// create builder to read into from the markdown file and parse line-based formatting (lists, headings)
	var builder strings.Builder
	builder.WriteString(`{{ define "content" }}`)

	heading := regexp.MustCompile(`(^#{0,6}\s.)`)
	list := regexp.MustCompile(`^(\s+(-|\+|\*|([a-z]\.)|([\d+]\.)|([i|v|x|c]+\.)))|(-\s)`)

	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())

		// check for line-specific formatting characters
		var formatted string
		if len(line) > 0 {
			if heading.FindIndex(line) != nil {
				formatted = handleHeadings(line, scanner.Text(), post)
			} else if list.FindIndex(line) != nil {
				formatted = handleList(line, scanner.Text())
			} else {
				formatted = scanner.Text()
			}
			builder.WriteString(formatted)
			builder.WriteString("\n")
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
	buf = cleanup(buf)
	return buf
}

func cleanup(content string) string {
	cleaned := strings.ReplaceAll(content, `\*`, "*")
	cleaned = strings.ReplaceAll(cleaned, `\~`, "~")
	return cleaned
}
