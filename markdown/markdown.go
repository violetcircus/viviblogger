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

	// fmt.Println("line:", string(line))

	// handle titles
	if line[0] == '#' {
		count := bytes.Count(line, []byte("#"))
		switch count {
		case 1:
			headerReplacer(content, 1, &builder)
		case 2:
			headerReplacer(content, 2, &builder)
		case 3:
			headerReplacer(content, 3, &builder)
		case 4:
			headerReplacer(content, 4, &builder)
		case 5:
			headerReplacer(content, 5, &builder)
		case 6:
			headerReplacer(content, 6, &builder)
		}
	}
	fmt.Println(builder.String())
}
