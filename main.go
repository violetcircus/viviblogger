package main

import (
	"bufio"
	// "fmt"
	"github.com/violetcircus/viviblogger/markdown"
	"github.com/violetcircus/viviblogger/output"
	"log"
	"os"
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

	var post output.Post
	scanner := bufio.NewScanner(file)
	post.Body = markdown.Convert(scanner, &post)

	output.Build(post)
}
