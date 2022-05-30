package main

import (
	"fmt"
	"log"
	"strconv"
)

func BinarizeString(n int) string {
	var s string
	for _, c := range fmt.Sprint(n) {
		s = fmt.Sprintf("%s%b", s, c)
	}
	return s
}

func ConcatenatedBinary(n int) int {
	binarized := BinarizeString(n)
	decimal, _ := strconv.ParseInt(binarized, 16, 0)
	log.Printf("%d", decimal)
	return int(decimal)
}
