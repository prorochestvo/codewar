package codewar

// https://www.codewars.com/kata/514b92a657cdc65150000006/train/go
func Multiple3And5(number int) int {
	sum := 0
	for i := 1; ; i++ {
		if v := 3 * i; v < number {
			sum += v
		} else {
			break
		}
		if v := 5 * i; v < number {
			sum += v
		}
	}
	return sum
}
