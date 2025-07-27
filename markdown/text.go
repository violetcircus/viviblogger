package markdown

//handle text decorations like bold, underline, italics etc. wrap anything surrounded by two newlines in a <p> tag i suppose

import (
	"fmt"
	"regexp"
	"strings"
)

// bold: [not an asterisk]**[text]**[not an asterisk]
// strikethrough is ~~[text]~~
// bold AND italic is ***[text]***
// will just use tags manually for underlining

// define replacer function
// match = found string
// tag = html tag string
// char = markdown format string
// trim = amount to trim the body of the text by - varies with number of characters used in md formatting string
func replacer(match string, tag string, char string, trim []int) string {
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
	opener := fmt.Sprintf("<%s>", tag)
	closer := fmt.Sprintf("</%s>", tag)

	// replace the first instance of char with <tag> and second with </tag>
	replaced := body[:start] + opener + body[start+trim[0]:end+trim[1]] + closer + body[end+1:]
	fmt.Println("match:", match, "result:", prefix+replaced)
	return prefix + replaced
}

// handles text. runs a find and replace over the entire html/markdown string as it exists so far
func handleText(content string) string {
	// regexes
	italics := regexp.MustCompile(`(?:^|[^\\*])\*([^\s*]+?[^\s])\*`)
	bold := regexp.MustCompile(`(?:^|[^\\*])\*\*([^\s\*][^*]+?[^\s])\*\*`)

	result := italics.ReplaceAllStringFunc(content, func(match string) string {
		return replacer(match, "i", "*", []int{1, 0})
	})
	result = bold.ReplaceAllStringFunc(result, func(match string) string {
		return replacer(match, "b", "*", []int{2, -1})
	})

	fmt.Println(result)
	return result
}
