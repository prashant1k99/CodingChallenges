package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func getBytesOfFile(fileContentInByte []byte) int {
	return len(fileContentInByte)
}

func readFileAndReturnContent(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []byte(""), err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}
	return content, nil
}

func main() {
	byteFlag := flag.Bool("c", false, "Get number of bytes for a file")
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		fmt.Println("No file name passed to process")
	}
	content, err := readFileAndReturnContent(args[0])
	if err != nil {
		fmt.Println("Unable to read file content")
		return
	}

	finalResponse := ""

	if *byteFlag {
		length := getBytesOfFile(content)
		finalResponse += strconv.Itoa(length) + "  "
	}
	finalResponse += args[0]
	fmt.Println(finalResponse)
}
