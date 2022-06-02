package main

func ValidateParentheses(s string) bool {
	opened := 0
	closed := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			opened++
		case ')':
			closed++
		}
	}
	return opened == closed
}
