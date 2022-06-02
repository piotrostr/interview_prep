# Generate Parentheses

given `n int` generate valid parentheses

Input: n = 1

`()` is good
`)(` is wrong

## Constraints

My understanding is that you need to construct a solution that will satisfy:

> open == closed == n

there are no more open parens than closed parens and they equal the desired
count

It can be constructed by going through all of the occurences while holding onto
two constraints:

> open > closed

if you run into `())` is never gonna be a valid solution

this makes it so that every time an open parentheses has to go before the closed
one

> open < n

one cannot generate more parentheses than n

## Implementation

Using backtracking here makes it so that you open a new function for every
condition and each function in go has a separate namespace, meaning that each
time you call open+1, this creates a new function with a namespace where the
open is 1 greater than the previous one.

The fail condition occurs naturally, since the function doesnt return after
going into the second if block (open < n), if after opening a new parentheses
there is not more open braces than closed one in the third if block (open <
closed) the function just finishes.
