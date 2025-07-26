package main

import (
	"bufio"
	"github.com/violetcircus/viviblogger/markdown"
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// var lines []string

	for scanner.Scan() {
		// lines = append(lines, scanner.Text())
		markdown.Convert(scanner.Text(), scanner)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal("scanner error:", err)
	}

	// for _, line := range lines {
	// 	fmt.Println(line)
	// }
}
