package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParenthesesCaseOne(t *testing.T) {
	assert.True(t, validateParenthesis("()"))
	assert.True(t, validateParenthesis("()"))
	assert.True(t, validateParenthesis("(()())"))
	assert.False(t, validateParenthesis("(()))"))
}

func TestGenerateParenthesesCaseOne(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	assert.Equal(t, want, generateParenthesis(3))
}

func TestGenerateParenthesesCaseTwo(t *testing.T) {
	want := []string{"()"}
	assert.Equal(t, want, generateParenthesis(1))
}

func TestGenerateParenthesesCaseThree(t *testing.T) {
	want := []string{"(())", "()()"}
	assert.Equal(t, want, generateParenthesis(2))
}

// thank god for github copilot
