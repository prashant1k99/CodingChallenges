package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func getBytesOfFile(fileContentInByte string) int {
	return len(fileContentInByte)
}

func getLineCountOfFile(fileContent string) int {
	return len(strings.Split(fileContent, "\n"))
}

func getWordCountOfFile(fileContent string) int {
	return len(strings.Fields(fileContent))
}

func getMultiByteCountForFile(fileContent string) int {
	return utf8.RuneCountInString(fileContent)
}

func readFileContentInBuffer(fileName string) (string, error) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	byteFlag := flag.Bool("c", false, "Get number of bytes for a file")
	lineFlag := flag.Bool("l", false, "Get number of lines for a file")
	wordFlag := flag.Bool("w", false, "Get number of words for a file")
	multiByteFlag := flag.Bool("m", false, "Get number of MultiByte character for a file")

	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		fmt.Println("No file name passed to process")
	}
	content, err := readFileContentInBuffer(args[0])
	if err != nil {
		fmt.Println("Unable to read file content")
		return
	}

	finalResponse := ""

	if *byteFlag {
		length := getBytesOfFile(content)
		finalResponse += strconv.Itoa(length) + "  "
	}
	if *lineFlag {
		length := getLineCountOfFile(content)
		finalResponse += strconv.Itoa(length) + "  "
	}
	if *wordFlag {
		length := getWordCountOfFile(content)
		finalResponse += strconv.Itoa(length) + "  "
	}
	if *multiByteFlag {
		length := getMultiByteCountForFile(content)
		finalResponse += strconv.Itoa(length) + "  "
	}
	finalResponse += args[0]
	fmt.Println(finalResponse)
}
