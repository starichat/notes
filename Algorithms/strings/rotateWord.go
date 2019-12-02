package main

import "fmt"
// 先将整个字符串逆序，然后将每个单词逆序
func rotateWord(str string) {

	// 保存每个单词
	var words []string
	reverse(str)
	// 用空格分隔字符串，获取每个自字符串

	if (str == nil ) {
		return
	}
}

func reverse(str string) string{
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	a := "dog is pig"
	fmt.Println(a)
	fmt.Println(reverse(a))
}