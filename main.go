package main

import (
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
		post := markdown.Read(args[1])
		output.Build(post)
	}
}
