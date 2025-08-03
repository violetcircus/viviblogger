package main

import (
	"github.com/violetcircus/viviblogger/configReader"
	"github.com/violetcircus/viviblogger/htmlWriter"
	"github.com/violetcircus/viviblogger/markdown"
	"log"
	"os"
)

func main() {
	configReader.MakeConfig()
	if len(os.Args) < 2 {
		log.Fatal("use the target filename as an argument numbnuts")
	} else {
		post := markdown.Read(os.Args[1])
		htmlWriter.Build(post)
	}
}
