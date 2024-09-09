package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type TokenType int

const (
	BraceOpen TokenType = iota
	BraceClose
	BracketOpen
	BracketClose
	Colon
	Comma
	Number
	String
	Bool
	Null
)

type Token struct {
	Value string
	Type  TokenType
}

func Tokenizer(contentToBeParsed string) ([]Token, error) {
	if len(contentToBeParsed) == 0 {
		return nil, fmt.Errorf("invalid input")
	}
	var tokenizedResponse []Token
	i := 0
	for i < len(contentToBeParsed) {
		char := contentToBeParsed[i]
		switch char {
		case '{':
			tokenizedResponse = append(tokenizedResponse, Token{
				string(char),
				BraceOpen,
			})
			i++
		case '}':
			tokenizedResponse = append(tokenizedResponse, Token{
				string(char),
				BraceClose,
			})
			i++
		case '[':
			tokenizedResponse = append(tokenizedResponse, Token{
				string(char),
				BracketOpen,
			})
			i++
		case ']':
			tokenizedResponse = append(tokenizedResponse, Token{
				string(char),
				BracketClose,
			})
			i++
		case ':':
			tokenizedResponse = append(tokenizedResponse, Token{
				string(char),
				Colon,
			})
			i++
		case ',':
			tokenizedResponse = append(tokenizedResponse, Token{
				string(char),
				Comma,
			})
			i++
		case '"':
			completeString, len := readCompleteString(contentToBeParsed[i:])
			tokenizedResponse = append(tokenizedResponse, Token{
				completeString,
				String,
			})
			i += len
		default:
			if unicode.IsSpace(rune(char)) {
				i++
			} else {
				return nil, fmt.Errorf("unable to parse content at [%d]", i)
			}
		}
	}
	return tokenizedResponse, nil
}

func readCompleteString(input string) (string, int) {
	i := 1
	for i < len(input) && input[i] != '"' {
		i++
	}
	return input[1:i], i + 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Unable to get the Stdin, Err:", err)
		os.Exit(1)
	}

	token, err := Tokenizer(string(content))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, t := range token {
		fmt.Println(t)
	}
	os.Exit(0)
}
