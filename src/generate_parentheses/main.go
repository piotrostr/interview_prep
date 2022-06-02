package main

func generateParentheses(n int) []string {
	open := 0
	closed := 0
	res := []string{}

	generate(n, open, closed, "", &res)
	return res
}

/*
My understanding is that you need to construct a solution that will satisfy:

  [x] open == closed == n

  there are no more open parens than closed parens and they equal the desired
  count

It can be constructed by going through all of the occurences while holding onto
two constraints:

 1) open > closed
         if you run into ()) is never gonna be a valid solution this
         makes it so that every time an open parentheses has to go before
         the closed one

 2) open < n
         one cannot generate more parentheses than n

Using backtracking here makes it so that you open a new function for every
condition and each function in go has a separate namespace, meaning that each
time you call open+1, this creates a new function with a namespace where the
open is 1 greater than the previous one.

The fail condition occurs naturally, since the function doesnt return after
going into the second if block (open < n), if after opening a new parentheses
there is not more open braces than closed one in the third if block (open <
closed) the function just finishes
*/
func generate(n int, open int, closed int, s string, res *[]string) {
	if open == n && closed == n {
		*res = append(*res, s)
		return
	}

	if open < n {
		generate(n, open+1, closed, s+"(", res)
	}

	if open > closed {
		generate(n, open, closed+1, s+")", res)
	}
}
