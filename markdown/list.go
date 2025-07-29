package markdown

// handle lists

import (
	// "bytes"
	"fmt"
	// "log"
	"regexp"
	// "slices"
	"strings"
)

// struct containing regex strings for formatting the lists
type listChecks struct {
	list *regexp.Regexp
	ul   *regexp.Regexp
	ol   *regexp.Regexp
}

// type that contains a list-closing HTML tag and how indented the line it came from was.
type listPosition struct {
	level int
	tag   string
}

func handleList(buf string) string {
	content := strings.Split(buf, "\n")

	checks := listChecks{
		list: regexp.MustCompile(`^(\s*)([-+*]|\d+\.|[a-z]\.|[ivxc]+\.)\s`),
		ul:   regexp.MustCompile(`^(?:\s*)([-+*]+)\s`),
		ol:   regexp.MustCompile(`^(?:\s*)(\d+\.|[a-z]\.|[ivxc]+\.)\s`),
	}

	stack := &[]listPosition{}
	var formatted []string
	for i, line := range content {
		if checks.list.MatchString(line) {
			formatted = append(formatted, listFormat(i, content, checks, stack))
		} else {
			formatted = append(formatted, line)
		}
	}

	fmt.Print(strings.Join(formatted, "\n"))
	return strings.Join(formatted, "\n")
}

func listFormat(i int, content []string, checks listChecks, stack *[]listPosition) string {
	var builder strings.Builder
	// get current, previous and next lines
	line := content[i]
	var prev string
	if i == 0 {
		prev = " "
	} else {
		prev = content[i-1]
	}
	var next string
	if i < len(content)-1 {
		next = content[i+1]
	} else {
		next = ""
	}

	// get whitespace count of previous, current and next lines
	prevCount := len(prev) - len(strings.TrimLeft(prev, " "))
	count := len(line) - len(strings.TrimLeft(line, " "))
	nextCount := len(next) - len(strings.TrimLeft(next, " "))

	// check whether the previous and next lines are lists
	prevList := checks.list.MatchString(prev)
	nextList := checks.list.MatchString(next)

	// check whether the next line is the opposite type of list, whether the previous line was, and also get the correct tag to use
	swapNext := false
	swapPrev := false
	var tag string
	if checks.ul.MatchString(line) {
		tag = "u"
		if checks.ol.MatchString(next) {
			swapNext = true
		}
		if checks.ol.MatchString(prev) {
			swapPrev = true
		}
	} else if checks.ol.MatchString(line) {
		tag = "o"
		if checks.ul.MatchString(next) {
			swapNext = true
		}
		if checks.ul.MatchString(prev) {
			swapPrev = true
		}
	}
	opener := fmt.Sprintf("<%sl>", tag)
	closer := fmt.Sprintf("</%sl>", tag)

	// replace markdown list formatting with a list item tag
	line = checks.list.ReplaceAllString(line, "<li>")

	// if the previous line isnt a list, or the current line is more indented than the last one, or the previous line WAS a list, AND it was a different kind from the current type of list, AND the whitespace counts are the same, add an opening tag.
	if !prevList || count > prevCount || prevList && swapPrev && count == prevCount {
		*stack = append(*stack, listPosition{count, closer})
		builder.WriteString(opener)
		builder.WriteRune('\n')
		// if the current line is less indented than the previous, start closing tags
	} else if count < prevCount {
		// loop backwards through the stack popping the tags with higher indent levels than the current line
		for i := len(*stack) - 1; i >= 0; i-- {
			item := (*stack)[i]
			if item.level > count {
				builder.WriteString(item.tag)
				fmt.Println("ITEM TAG!!!", item.tag)
				*stack = (*stack)[:len(*stack)-1]
			}
		}
	}

	// write the content of the line finally
	builder.WriteString(line)
	builder.WriteString("</li>")

	// if the next line isnt a list, OR the whitespace amounts on this line and the next are the same AND it's a new type of list:
	if !nextList || nextCount == count && swapNext == true {
		builder.WriteRune('\n')
		builder.WriteString(closer)
	}

	result := builder.String()
	return result
}
