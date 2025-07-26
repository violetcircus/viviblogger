package markdown

import (
	"bytes"
	"fmt"
	"github.com/violetcircus/viviblogger/output"
	"strings"
)

func handleHeadings(line []byte, content string, builder *strings.Builder, post *output.Post) {
	if line[0] == '#' {
		// count number of #s at beginning of line
		count := bytes.Count(line, []byte("#"))
		str := strings.Replace(content, "#", "", -1)
		// if the post doesnt have a title already, set it to the first h1.
		if count == 1 && post.Title == "" {
			post.Title = str
		} else {
			builder.WriteString(fmt.Sprintf("<h%d>%s</h%d>", count, strings.TrimSpace(str), count))
		}
	}
}
