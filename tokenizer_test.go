package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// combination  expression: (add 2 (subtract 4 2))
// single expression: (add 1 2)
// single number: 2
// single paren: (  / )

// func Test
func TestCombinationExpression(t *testing.T) {
	t.Skip(".")
}

func TestSingleExpression(t *testing.T) {
	t.Skip(".")
}

func TestParen(t *testing.T) {
	tokens := []token{
		{
			kind:  "Paren",
			value: "(",
		},
	}

	code := "("
	assert.Equal(t, tokens, tokenizer(code))
}

func tokenizer(code string) []token {
	tokens := []token{}
	current := 0
	for current < len(code) {
		char := string(code[current])

		if char == "(" {
			tokens = append(tokens, token{
				kind:  "Paren",
				value: "(",
			})
			current++
		}

	}
	return tokens
}

func TestNumber(t *testing.T) {
	t.Skip(".")
}

type token struct {
	kind  string
	value string
}
