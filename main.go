package main

import (
	"github.com/violetcircus/viviblogger/markdown"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(os.Args) < 2 {
		log.Fatal("use the target filename as an argument numbnuts")
	} else {
		markdown.Read(args[1])
	}
}
