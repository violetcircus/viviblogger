package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(os.Args) < 2 {
		log.Fatal("use the target filename as an argument numbnuts")
	} else {
		Read(args[1])
	}
}

func Read(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open %s, %v", fileName, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// var lines []string

	for scanner.Scan() {
		// lines = append(lines, scanner.Text())
		Convert(scanner.Text(), scanner)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal("scanner error:", err)
	}

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
}

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
		}
		fmt.Println(builder.String())
	}
}
func headerReplacer(line string, num int, builder *strings.Builder) {
	str := strings.Replace(line, "#", "", -1)
	builder.WriteString(fmt.Sprintf("<h%d>%s</h%d>", num, strings.TrimSpace(str), num))
}
