package main

func generateParentheses(n int) []string {
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
