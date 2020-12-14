package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	Names = make([]string, 1)
	FailedList = make([]bool, 1)
)

func PrintList() {
	realStudentSize := 0

	for index, val := range Names {
		if val != "" {
			realStudentSize++
			compile := "success"
			if FailedList[index] {
				compile = "failed"
			}

			fmt.Printf("filename: %s, compile: %s\n", val, compile)
		}
	}

	fmt.Printf("total student: %d", realStudentSize)
}

func CheckError(err error, msg string) bool {
	if err != nil {
		log.Println(msg)
		log.Fatal(err)
		return false
	}

	return true
}

func ReadFiles(dir string, target string) {
	i := 0
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			name := info.Name()
			if strings.Contains(name, target) {
				// 파일 이름 추가
				Names[i] = info.Name()
				i++
				Evaluate(path, "", i)
			}
			return nil
		})

	CheckError(err, "err occurred while reading directories")
}

func CreateTempFolder() error {
	dir, _ := os.Getwd()

	if _, err := os.Stat(dir + "/tmp"); os.IsNotExist(err) {
		err = os.Mkdir(dir + "/tmp", os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func Evaluate(path string, input string, index int) {
	// create folder
	err := CreateTempFolder()
	CheckError(err, "failed to create temp folder")

	// compile
	dir, _ := os.Getwd()
	cmd := exec.Command("gcc", path, "-o", dir + "/tmp/main")
	err = cmd.Run()
	if err != nil  {
		FailedList[index] = true
		return
	}
}

func main() {
	// command line flags
	dir := flag.String("d", "/Users/idohyeon/Downloads/practices/01", "directory")
	studentSize := flag.String("s", "50", "student size")
	inputFile := flag.String("i", "", "input file")
	outputFile := flag.String("o", "", "output file")
	targetFiles := flag.String("t", "p1.c", "target files format")
	flag.Parse()

	size, err := strconv.Atoi(*studentSize)
	CheckError(err, "student size receives only numbers.")

	Names = make([]string, size)
	FailedList = make([]bool, size)

	fmt.Printf("Received dir: %q, input: %q, output: %q, targetFile: %q \n", *dir, *inputFile, *outputFile, *targetFiles)

	ReadFiles(*dir, *targetFiles)

	PrintList()
}