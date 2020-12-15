package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
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
	RuntimeErrorList = make([]bool, 1)
	OutputList = make([]string, 1)
)

func PrintList() {
	realStudentSize := 0
	bluePrint := color.New(color.FgBlue).PrintfFunc()
	greenPrint := color.New(color.FgGreen).PrintfFunc()

	for index, val := range Names {
		if val != "" {
			compile := "success"
			if FailedList[index] {
				compile = "failed"
			}

			runtime := "success"
			if RuntimeErrorList[index] {
				runtime = "failed"
			}

			fmt.Printf("filename: %s, compile: %s, runtime: %s\n", val, compile, runtime)
			realStudentSize++
		}
	}

	bluePrint("total student: %d\n\n", realStudentSize)
	greenPrint("---- output ----\n\n")

	realStudentSize = 0
	for index, val := range Names {
		if val != "" {
			if !RuntimeErrorList[index] && !FailedList[index] {
				greenPrint("[ filename: %s ]\n", val)
				fmt.Printf("%s\n\n", OutputList[index])
				realStudentSize++
			}
		}
	}
}

func CheckError(err error, msg string) bool {
	if err != nil {
		log.Println(msg)
		log.Fatal(err)
		return false
	}

	return true
}

func ReadFiles(dir string, target string, inputFile string) {
	i := 0
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			name := info.Name()
			if strings.Contains(name, target) {
				// 파일 이름 추가
				Names[i] = info.Name()
				Evaluate(path, inputFile, i)
				i++
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
		RuntimeErrorList[index] = true
		return
	}

	if input == "" {
		cmd = exec.Command(dir + "/tmp/main")
	} else {
		cmd = exec.Command(dir + "/tmp/main", "<", input)
	}

	out, err := cmd.Output()
	if err != nil {
		RuntimeErrorList[index] = true
		return
	}

	OutputList[index] = string(out)
}

func main() {
	// command line flags
	dir := flag.String("d", "", "directory")
	studentSize := flag.String("s", "50", "student size")
	inputFile := flag.String("i", "", "input file")
	targetFiles := flag.String("t", "", "target files format")
	flag.Parse()

	size, err := strconv.Atoi(*studentSize)
	CheckError(err, "student size receives only numbers.")

	if *dir == "" && *targetFiles == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	notiPrint := color.New(color.FgWhite).Add(color.BgGreen).PrintfFunc()
	Names = make([]string, size)
	FailedList = make([]bool, size)
	RuntimeErrorList = make([]bool, size)
	OutputList = make([]string, size)

	fmt.Printf("dir: ")
	notiPrint("%q\n", *dir)

	fmt.Printf("students: ")
	notiPrint("%q\n", *studentSize)

	fmt.Printf("input: ")
	notiPrint("%q\n", *inputFile)

	fmt.Printf("target: ")
	notiPrint("%q\n\n", *targetFiles)

	ReadFiles(*dir, *targetFiles, *inputFile)

	PrintList()
}