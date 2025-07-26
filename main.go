package main

import (
	"bufio"
	"fmt"
	"github.com/violetcircus/viviblogger/markdown"
	"github.com/violetcircus/viviblogger/output"
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

	var p output.Post
	p.Title = ""

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf(`{{ define "content" }}`))
	for scanner.Scan() {
		markdown.Convert(scanner.Text(), scanner, &builder, &p)
	}
	builder.WriteString(`{{ end }}`)

	err = scanner.Err()
	if err != nil {
		log.Fatal("scanner error:", err)
	}

	p.Preview = "idk yet lmaoo"
	p.Body = builder.String()
	output.Build(p)
}
