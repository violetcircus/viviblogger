package markdown

import (
	"bytes"
	"fmt"
	"github.com/violetcircus/viviblogger/htmlWriter"
	"strings"
)

func handleHeadings(line []byte, content string, post *htmlWriter.Post) string {
	// count number of #s at beginning of line
	count := bytes.Count(line, []byte("#"))
	str := strings.Replace(content, "#", "", -1)
	// if the post doesnt have a title already, set it to the first h1.
	if count == 1 && post.Title == "" {
		post.Title = str
	} else {
		return fmt.Sprintf("<h%d>%s</h%d>", count, strings.TrimSpace(str), count)
	}
	return ""
}
