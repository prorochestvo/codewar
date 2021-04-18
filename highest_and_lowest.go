package codewar

import (
	"strconv"
	"strings"
)

// https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go
func HighAndLow(in string) string {
	a := strings.Split(in, " ")
	max := 0
	min := 0
	for i, s := range a {
		n, _ := strconv.Atoi(s)
		if i == 0 || n < min {
			min = n
		}
		if i == 0 || n > max {
			max = n
		}
	}
	return strconv.Itoa(max) + " " + strconv.Itoa(min)
}
