package markdown

// handle lists

import (
	// "bytes"
	"log"
	"regexp"
	"strings"
)

func handleList(buf string) string {
	// line []byte, prev string, content string
	content := strings.Split(buf, "\n")

	ul := regexp.MustCompile(`^(\s*)([-+*]+)\s`)
	ol := regexp.MustCompile(`^(\s*)(\d+\.|[a-z]\.|[ivxc]+\.)\s`)

	var formatted []string
	for i, line := range content {
		if ul.MatchString(line) {
			formatted = append(formatted, unorderedList(i, content, ul))
		} else if !ol.MatchString(line) {
			formatted = append(formatted, orderedList(i, content, ul))
		} else {
			formatted = append(formatted, line)
		}
	}

	var result []string
	result = append(result, content[0])
	result = append(result, strings.Join(formatted, "\n"))
	result = append(result, content[len(content)-1])

	return strings.Join(result, "\n")
}

func unorderedList(i int, content []string, ul *regexp.Regexp) string {
	line := content[i]
	var prev string
	if i == 0 {
		prev = " "
	} else {
		prev = content[i-1]
	}
	var next string
	if i < len(content)-1 {
		next = content[i+1]
	} else {
		next = ""
	}

	prevCount := len(prev) - len(strings.TrimLeft(prev, " "))
	count := len(line) - len(strings.TrimLeft(line, " "))

	var builder strings.Builder

	line = ul.ReplaceAllString(line, "<li>")

	prevList := ul.MatchString(prev)
	nextList := ul.MatchString(next)

	log.Println(prevCount, prev, "prevList:", prevList, nextList)

	if !prevList || count > prevCount {
		builder.WriteString("<ul>")
		builder.WriteRune('\n')
	} else if count < prevCount && prevCount != 0 {
		builder.WriteString("</ul>")
		builder.WriteRune('\n')
	}

	builder.WriteString(line)
	builder.WriteString("</li>")
	if !nextList {
		builder.WriteRune('\n')
		builder.WriteString("</ul>")
	}

	return builder.String()
}

func orderedList(i int, content []string, ol *regexp.Regexp) string {

	return ""
}
