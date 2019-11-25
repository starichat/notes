package main

import "fmt"

func isRotate(s, d string) bool {
	s2 := d + d
	if -1 == getIndexOf(s2, s) {
		return false
	}
	return true
}

func getIndexOf(s string, pattern string) int {

	n := len(s)
	m := len(pattern)
	if n < m {
		return -1
	}
	nexts := getNexts(pattern)

	j := 0
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != pattern[j] {
			j = nexts[j-1] + 1
		}

		if s[i] == pattern[j] {
			if j == m-1 {
				return i - m + 1
			}
			j += 1
		}
	}
	return -1
}

func getNexts(pattern string) []int {
	m := len(pattern)
	nexts := make([]int, m)
	for index := range nexts {
		nexts[index] = -1
	}

	for i := 1; i < m-1; i++ {
		j := nexts[i-1]

		for pattern[j+1] == pattern[i] && j >= 0 {
			j = nexts[j]
		}

		nexts[i] = j
	}
	return nexts
}
func main() {
	fmt.Println(isRotate("abcd", "abcd"))
	fmt.Println(isRotate("abcde", "eabcd"))
}
