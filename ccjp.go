package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
			} else if unicode.IsDigit(rune(char)) || char == '-' {
				numericValue, len := readCompleteNumber(contentToBeParsed[i:])
				tokenizedResponse = append(tokenizedResponse, Token{
					numericValue,
					Number,
				})
				i += len
				// Get the number content as the input
			} else if strings.HasPrefix(contentToBeParsed[i:], "true") {
				tokenizedResponse = append(tokenizedResponse, Token{
					"true",
					Bool,
				})
				i += 4
			} else if strings.HasPrefix(contentToBeParsed[i:], "false") {
				tokenizedResponse = append(tokenizedResponse, Token{
					"false",
					Bool,
				})
				i += 5
			} else if strings.HasPrefix(contentToBeParsed[i:], "null") {
				tokenizedResponse = append(tokenizedResponse, Token{
					"null",
					Null,
				})
				i += 4
			} else {
				return nil, fmt.Errorf("unable to parse content at [%d]", i)
			}
		}
	}
	return tokenizedResponse, nil
}

func readCompleteNumber(input string) (string, int) {
	i := 0
	for i < len(input) && (unicode.IsDigit(rune(input[i])) || input[i] == '.' || input[i] == '-') {
		i++
	}
	return input[:i], i
}

func readCompleteString(input string) (string, int) {
	i := 1
	for i < len(input) && input[i] != '"' {
		i++
	}
	return input[1:i], i + 1
}

type (
	JSONObject map[string]interface{}
	JSONArray  []interface{}
)

func parseArray(tokens []Token) (JSONArray, error) {
	return nil, nil
}

func parseObject(tokens []Token) (JSONObject, error) {
	return nil, nil
}

func JSONParser(tokens []Token) (interface{}, error) {
	initialToken := tokens[0]

	switch initialToken.Type {
	case BraceOpen:
		parsedObject, err := parseObject(tokens)
		if err != nil {
			return nil, err
		}
		return parsedObject, nil
	case BracketOpen:
		parsedArray, err := parseArray(tokens)
		if err != nil {
			return nil, err
		}
		return parsedArray, nil
	default:
		return nil, fmt.Errorf("invalid value at index 0")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Unable to get the Stdin, Err:", err)
		os.Exit(1)
	}

	tokens, err := Tokenizer(string(content))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parsedInput, err := JSONParser(tokens)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(parsedInput)
	os.Exit(0)
}
