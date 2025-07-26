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
		switch count {
		case 1:
			headingReplacer(content, 1, builder)
		case 2:
			headingReplacer(content, 2, builder)
		case 3:
			headingReplacer(content, 3, builder)
		case 4:
			headingReplacer(content, 4, builder)
		case 5:
			headingReplacer(content, 5, builder)
		case 6:
			headingReplacer(content, 6, builder)
		}
	}
}

func headingReplacer(line string, num int, builder *strings.Builder) {
	str := strings.Replace(line, "#", "", -1)
	builder.WriteString(fmt.Sprintf("<h%d>%s</h%d>", num, strings.TrimSpace(str), num))
}
