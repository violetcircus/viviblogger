package markdown

// the main markdown package file

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func Convert(content string, scanner *bufio.Scanner) {
	var builder strings.Builder
	line := bytes.TrimSpace(scanner.Bytes())
	if len(line) > 0 {
		handleHeadings(line, content, &builder)
	}
	fmt.Println(builder.String())
}
