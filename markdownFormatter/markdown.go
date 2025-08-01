package markdownFormatter

// the main markdown package file

import (
	"bufio"
	"bytes"
	"github.com/violetcircus/viviblogger/htmlWriter"
	"log"
	"os"
	"regexp"
	"strings"
)

// post title will be first h1. somehow need to figure out what the preview is, too

func Read(fileName string) htmlWriter.Post {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open %s, %v", fileName, err)
	}
	defer file.Close()

	var post htmlWriter.Post
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	post.Body = convert(scanner, &post)
	return post
}

func convert(scanner *bufio.Scanner, post *htmlWriter.Post) string {
	// create builder to read into from the markdown file and parse line-based formatting (lists, headings)
	var builder strings.Builder
	builder.WriteString(`{{ define "content" }}`)

	heading := regexp.MustCompile(`(^#{0,6}\s.)`)

	// prev := make([]byte, 1024)
	for scanner.Scan() {
		log.Println("scanner text:", scanner.Text())
		line := bytes.TrimSpace(scanner.Bytes())

		// check for line-specific formatting characters
		var formatted string
		if len(line) > 0 {
			if heading.FindIndex(line) != nil {
				formatted = handleHeadings(line, scanner.Text(), post)
			} else {
				formatted = scanner.Text()
			}
			builder.WriteString(formatted)
			builder.WriteString("\n")
		} else {
			formatted = scanner.Text()
		}
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal("scanner error:", err)
	}
	// end the template
	builder.WriteString(`{{ end }}`)

	// start reformatting through replacing in entire string
	buf := builder.String()
	buf = handleText(buf)
	buf = handleList(buf)
	buf = handleParagraphs(buf)
	buf = handleLinks(buf)
	buf = cleanup(buf)
	log.Print(buf)
	return buf
}

func cleanup(content string) string {
	cleaned := strings.ReplaceAll(content, `\*`, "*")
	cleaned = strings.ReplaceAll(cleaned, `\~`, "~")
	return cleaned
}
