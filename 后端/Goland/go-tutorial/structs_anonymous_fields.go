package main

import "fmt"

/**
结构体可以包含一个或多个匿名字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型。
Go语言的集成是通过内嵌或组合来实现的，因此在Go语言中，继承相比于组合，组合更好
**/
type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

// 内嵌结构体方式实现继承
// type A struct {
// 	ax, ay int
// }

// type B struct {
// 	A
// 	bx, by float32
// }

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}
