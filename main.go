package main

import (
	"github.com/violetcircus/viviblogger/htmlWriter"
	"github.com/violetcircus/viviblogger/markdownFormatter"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("use the target filename as an argument numbnuts")
	} else {
		post := markdownFormatter.Read(os.Args[1])
		htmlWriter.Build(post)
	}
}
