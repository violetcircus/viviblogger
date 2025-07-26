package markdown

import (
	"bytes"
	"fmt"
	"strings"
)

func handleHeadings(line []byte, content string, builder *strings.Builder) {
	if line[0] == '#' {
		// count number of #s at beginning of line
		count := bytes.Count(line, []byte("#"))
		str := strings.Replace(content, "#", "", -1)

		builder.WriteString(fmt.Sprintf("<h%d>%s</h%d>", count, strings.TrimSpace(str), count))
	}
}
