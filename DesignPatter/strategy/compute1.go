package strategy

import "fmt"

type Multiplication struct{}

func (d Multiplication) Compute(num1, num2 int) int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
			return
		}
	}()

	if num2 == 0 {
		panic("num2 must not be zero!")
	}

	return num1 * num2
}
