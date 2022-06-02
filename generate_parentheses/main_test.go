package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParenthesesCaseOne(t *testing.T) {
	assert.True(t, ValidateParentheses("()"))
	assert.True(t, ValidateParentheses("()"))
	assert.True(t, ValidateParentheses("(()())"))
	assert.False(t, ValidateParentheses("(()))"))
}

func TestGenerateParenthesesCaseOne(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	assert.Equal(t, want, generateParentheses(3))
}

func TestGenerateParenthesesCaseTwo(t *testing.T) {
	want := []string{"()"}
	assert.Equal(t, want, generateParentheses(1))
}

func TestGenerateParenthesesCaseThree(t *testing.T) {
	want := []string{"(())", "()()"}
	assert.Equal(t, want, generateParentheses(2))
}

// thank god for github copilot
