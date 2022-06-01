package main

// lets start by outlining the algorithm in pseudocode
// start by creating a full parentheses
// keep track of the number of complete parentheses and decrement n

// lets start with a brute force approach

func validateParenthesis(s string) bool {
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

func generateParenthesis(n int) []string {
	open := 0
	closed := 0
	res := []string{}

	generate(n, open, closed, "", &res)
	return res
}

func generate(n int, open int, closed int, s string, res *[]string) {
	if open < n {
		generate(n, open+1, closed, s+"(", res)
	}

	if open > closed {
		generate(n, open, closed+1, s+")", res)
	}

	if open == n && closed == n {
		*res = append(*res, s)
	}
}
