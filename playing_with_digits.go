package codewar

import (
	"strconv"
)

func pow(x, y int) int {
	if y == 0 || x == 1 {
		return 1
	}
	return x * pow(x, y-1)
}

// https://www.codewars.com/kata/5552101f47fc5178b1000050/train/go
func DigPow(n, p int) int {
	v := strconv.Itoa(n)
	l := len(v)
	sum := 0
	for i := 0; i < l; i++ {
		n := int(v[i]) - 48
		sum += pow(n, p)
		p++
	}
	if sum%n != 0 {
		return -1
	}
	return sum / n
}
