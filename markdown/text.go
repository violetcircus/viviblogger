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

func handleText(content string) string {
	italics := regexp.MustCompile(`(?:^|[^\\*])\*([^*]+?)\*`)
	result := italics.ReplaceAllStringFunc(content, func(match string) string {
		// filters out the extra char before the first * (go uses old regex)
		prefix := match[:1]
		body := match[1:]

		// replace the first * with <i> and second * with </i>
		start := strings.Index(body, "*")
		end := strings.LastIndex(body, "*")

		if start == -1 || end == -1 || start == end {
			return match
		}

		replaced := body[:start] + "<i>" + body[start+1:end] + "</i>" + body[end+1:]
		return prefix + replaced
	})
	fmt.Println(result)
	return result
}
