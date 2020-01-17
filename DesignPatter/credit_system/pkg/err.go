package main

import (
    "fmt"
)

func divide(a, b int) int {
    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("Runtime panic caught: %v\n", err)
        }
    }()


    k := a / b
	fmt.Printf("%d / %d = %d\n", a, b, k)
	return k
}

func main() {
    fmt.Println(divide(1,0))
    fmt.Println("divide?????????main??")
}
