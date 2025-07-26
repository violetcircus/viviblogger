package markdown

// handle headings
import (
	"fmt"
	"strings"
)

func headerReplacer(line string, num int, builder *strings.Builder) {
	str := strings.Replace(line, "#", "", -1)
	builder.WriteString(fmt.Sprintf("<h%d>%s</h%d>", num, strings.TrimSpace(str), num))
}
