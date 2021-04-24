package codewar

import "strings"

// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1/train/go
func CountingDuplicates(s string) (c int)  {
	s = strings.ToLower(s)
	for len(s) > 1 {
		if strings.Index(s[1:], s[0:1]) >= 0 {
			c++
		}
		s = strings.ReplaceAll(s, s[0:1], "")
	}
	return c
}

func CountingDuplicatesV2(s string) (c int) {
	mark := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if mark[r]++; mark[r] == 2 {
			c++
		}
	}
	return
}

func CountingDuplicatesV3(s string) (c int) {
	s = strings.ToLower(s)
	l := len(s)
	mark := make([]bool, l)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] && mark[j] == false {
        if mark[i] == false {
          c++
          mark[i] = true
        }
        mark[j] = true
			}
		}
	}
	return
}
