package markdown

// handle links and image embeds
import (
	"fmt"
	"regexp"
	"strings"
)

func handleLinks(content string) string {
	link := regexp.MustCompile(`[\[][^\\]*[\]]\(.*\)`)
	result := link.ReplaceAllStringFunc(content, func(match string) string {
		textStart := strings.IndexRune(match, '[')
		textEnd := strings.IndexRune(match, ']')
		text := match[textStart+1 : textEnd]

		urlStart := strings.IndexRune(match, '(')
		urlEnd := strings.IndexRune(match, ')')
		url := match[urlStart+1 : urlEnd]
		result := fmt.Sprintf("text: %s, url: %s", text, url)

		return result
	})
	return result
}
