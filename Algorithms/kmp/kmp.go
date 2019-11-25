package main

import "fmt"

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

	s := "adafafafafa  afaf afa"
	pattern := " afa"
	fmt.Println(getIndexOf(s, pattern))

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "abab"
	fmt.Println(getIndexOf(s, pattern)) //11

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "ababacd"
	fmt.Println(getIndexOf(s, pattern)) //-1

	s = "hello"
	pattern = "ll"
	fmt.Println(getIndexOf(s, pattern)) //2

}
