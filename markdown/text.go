package markdown

//handle text decorations like bold, strikethrough, italics etc. wrap anything surrounded by two newlines in a <p> tag i suppose

import (
	"fmt"
	"regexp"
	"strings"
)

// strikethrough is ~~[text]~~

// define replacer function
// match = found string
// tag = html tag string
// char = markdown format string
// trim = amount to trim the body of the text by - varies with number of characters used in md formatting string
func replacer(match string, opener string, closer string, char string, trim []int) string {
	// filters out the extra char before the first * (go uses an older regex thing. had to do extra filtering)
	prefix := match[:1] // extra char
	body := match[1:]   // rest of string

	start := strings.Index(body, char)   // first char index
	end := strings.LastIndex(body, char) // last char index

	// fmt.Println(tag, "match", match)
	// fmt.Println(tag, "prefix", prefix)
	// fmt.Println(tag, "body", body)

	// if theres no asterisks then return string as-is
	if start == -1 || end == -1 || start == end {
		fmt.Println("failed!")
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
		return replacer(match, "<i>", "</i>", "*", []int{1, 0})
	})
	result = bold.ReplaceAllStringFunc(result, func(match string) string {
		return replacer(match, "<b>", "</b>", "*", []int{2, -1})
	})
	result = boldtalic.ReplaceAllStringFunc(result, func(match string) string {
		return replacer(match, "<b><i>", "</b></i>", "*", []int{3, -2})
	})
	result = strikethrough.ReplaceAllStringFunc(result, func(match string) string {
		return replacer(match, "<s>", "</s>", "~", []int{2, -1})
	})

	fmt.Print(result)
	return result
}
