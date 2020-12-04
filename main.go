package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CheckError(err error, msg string) bool {
	if err != nil {
		log.Println(msg)
		log.Fatal(err)
		return false
	}

	return true
}

func ReadFiles(dir string) {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})

	CheckError(err, "err occurred while reading directories")
}

func main() {
	// command line flags
	dir := flag.String("d", "", "directory")
	inputFile := flag.String("i", "", "input file")
	outputFile := flag.String("o", "", "output file")
	targetFiles := flag.String("t", "", "target files format")
	flag.Parse()

	fmt.Printf("Received dir: %q, input: %q, output: %q, targetFile: %q \n", *dir, *inputFile, *outputFile, *targetFiles)
	fmt.Println("Hello, World!")

	ReadFiles(*dir)
}