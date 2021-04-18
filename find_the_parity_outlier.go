package codewar

// https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
func FindOutlier(integers []int) int {
	result := -1
	need := 0 // "need" is even integer
	if ((integers[0] & 1) + (integers[1] & 1) + (integers[2] & 1)) < 2 {
		need = 1 // "need" is odd integer
	}
	for _, num := range integers {
		// checking even or odd integer
		if num&1 == need {
			result = num
			break
		}
	}
	return result
}
