package main

import (
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
		hasError bool
	}{
		{
			input: `{"key": "value"}`,
			expected: []Token{
				{Value: "{", Type: BraceOpen},
				{Value: "key", Type: String},
				{Value: ":", Type: Colon},
				{Value: "value", Type: String},
				{Value: "}", Type: BraceClose},
			},
			hasError: false,
		},
		{
			input:    `{invalid}`,
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		tokens, err := Tokenizer(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Tokenizer() error = %v, wantErr %v", err, test.hasError)
			continue
		}
		if !reflect.DeepEqual(tokens, test.expected) {
			t.Errorf("Tokenizer() = %v, want %v", tokens, test.expected)
		}
	}
}

func TestJSONParser(t *testing.T) {
	tests := []struct {
		expected interface{}
		input    []Token
		hasError bool
	}{
		{
			input: []Token{
				{Value: "{", Type: BraceOpen},
				{Value: "key", Type: String},
				{Value: ":", Type: Colon},
				{Value: "value", Type: String},
				{Value: "}", Type: BraceClose},
			},
			expected: JSONObject{"key": "value"},
			hasError: false,
		},
		{
			input:    []Token{{Value: "{", Type: BraceOpen}},
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := JSONParser(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("JSONParser() error = %v, wantErr %v", err, test.hasError)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("JSONParser() = %v, want %v", result, test.expected)
		}
	}
}
