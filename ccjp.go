package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	for index, content := range contentToBeParsed {
		switch content {
		case '{':
			tokenizedResponse = append(tokenizedResponse, Token{
				Value: string(content),
				Type:  BraceOpen,
			})
		case '}':
			tokenizedResponse = append(tokenizedResponse, Token{
				Value: string(content),
				Type:  BraceClose,
			})
		case '[':
			tokenizedResponse = append(tokenizedResponse, Token{
				Value: string(content),
				Type:  BracketOpen,
			})
		case ']':
			tokenizedResponse = append(tokenizedResponse, Token{
				Value: string(content),
				Type:  BracketClose,
			})
		case ':':
			tokenizedResponse = append(tokenizedResponse, Token{
				Value: string(content),
				Type:  Colon,
			})
		default:
			return nil, fmt.Errorf("unable to parse content at [%d]", index)
		}
	}
	return tokenizedResponse, nil
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
