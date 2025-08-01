package markdownFormatter

//handle text decorations like bold, strikethrough, italics etc. also, paragraphs! was going to skip them, but I heard screen readers use them, so they can stay.

import (
	"log"
	"regexp"
	"strings"
)

// tagReplacer function, used in text.go
// match = found string
// tag = html tag string
// char = markdown format string
// trim = amount to trim the body of the text by - varies with number of characters used in md formatting string
func tagReplacer(match string, opener string, closer string, char string, trim []int) string {
	// filters out the extra char before the first * (go uses an older regex thing. had to do extra filtering)
	prefix := match[:1] // extra char
	body := match[1:]   // rest of string

	start := strings.Index(body, char)   // first char index
	end := strings.LastIndex(body, char) // last char index

	// if theres no asterisks then return string as-is
	if start == -1 || end == -1 || start == end {
		log.Println("failed!")
		return match
	}

	//special case for bold+italic
	replaced := body[:start] + opener + body[start+trim[0]:end+trim[1]] + closer + body[end+1:]
	return prefix + replaced
	// replace the first instance of char with <tag> and second with </tag>
}

// handles text. runs a find and replace over the entire html/markdown string as it exists so far
func handleText(content string) string {
	// regexes
	italics := regexp.MustCompile(`(?:^|[^\\*])\*([^\s^*].+?[^\s])\*[^\*]`)
	bold := regexp.MustCompile(`(?:^|[^\\*])\*\*([^\s\*].+?[^\s])\*\*`)
	boldtalic := regexp.MustCompile(`(?:^|[^\\*])\*\*\*([^\s\*].+?[^\s])\*\*\*`)
	strikethrough := regexp.MustCompile(`(?:^|[^\\~])\~\~([^\s\~].+?[^\s])\~\~`)

	result := italics.ReplaceAllStringFunc(content, func(match string) string {
		return tagReplacer(match, "<i>", "</i>", "*", []int{1, 0})
	})
	result = bold.ReplaceAllStringFunc(result, func(match string) string {
		return tagReplacer(match, "<b>", "</b>", "*", []int{2, -1})
	})
	result = boldtalic.ReplaceAllStringFunc(result, func(match string) string {
		return tagReplacer(match, "<b><i>", "</i></b>", "*", []int{3, -2})
	})
	result = strikethrough.ReplaceAllStringFunc(result, func(match string) string {
		return tagReplacer(match, "<s>", "</s>", "~", []int{2, -1})
	})

	return result
}

func handleParagraphs(buf string) string {
	content := strings.Split(buf, "\n")

	tag := regexp.MustCompile(`({{)|(^\s*$)|[<](.{0,2}l.{0,2}|[h]\d)[>]`)

	stack := []int{}
	formatted := []string{}
	for i, line := range content {
		var builder strings.Builder
		if len(stack) > 0 {
			builder.WriteString("</p>")
			builder.WriteString("\n")
			stack = stack[:len(stack)-1]
		}
		if !tag.MatchString(line) {
			builder.WriteString("<p>")
			builder.WriteString(line)
			if i < len(content) && tag.MatchString(content[i+1]) {
				builder.WriteString("</p>")
			} else {
				stack = append(stack, i)
			}
			formatted = append(formatted, builder.String())
		} else {
			builder.WriteString(line)
			formatted = append(formatted, builder.String())
		}
	}

	result := strings.Join(formatted, "\n")
	return result
}
