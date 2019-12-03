package main

import (
	"fmt"
	"strings"
)

// 先将整个字符串逆序，然后将每个单词逆序
func rotateWord(str string) string {

	// 保存每个单词

	// 用空格分隔字符串，获取每个自字符串
	var words []string
	fmt.Println(words)
	for _, s := range splitStr(reverse(str)) {
		words = append(words, reverse(s))

	}

	fmt.Println(words)
	return strings.Join(words, " ")

}

func reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func splitStr(str string) []string {
	return strings.Split(str, " ")
}

func main() {
	a := "dog is pig"
	//	fmt.Println(a)
	fmt.Println(rotateWord(a))
	//fmt.Println(reverse(a))
}
