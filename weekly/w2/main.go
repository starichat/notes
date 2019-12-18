package main

import "fmt"

//Person类型
type Person struct {
	name string
	age  int
}

//切片类型
type IntSlice []int
type FloatSlice []float32
type PersonSlice []Person

//接口定义
type MaxInterface interface {
	Len() int
	Get(i int) interface{}
	Bigger(i, j int) bool
}

//Len()方法的实现
func (x IntSlice) Len() int {
	return len(x)
}
func (x FloatSlice) Len() int {
	return len(x)
}
func (x PersonSlice) Len() int {
	return len(x)
}

//Get(i int)方法实现
func (x IntSlice) Get(i int) interface{} {
	return x[i]
}
func (x FloatSlice) Get(i int) interface{} {
	return x[i]
}
func (x PersonSlice) Get(i int) interface{} {
	return x[i]
}

//Bigger(i,j int)方法实现
func (x IntSlice) Bigger(i, j int) bool {
	if x[i] > x[j] {
		return true
	} else {
		return false
	}
}
func (x FloatSlice) Bigger(i, j int) bool {
	if x[i] > x[j] {
		return true
	} else {
		return false
	}
}
func (x PersonSlice) Bigger(i, j int) bool {
	if x[i].age > x[j].age {
		return true
	} else {
		return false
	}
}

//求最大值函数实现
func Max(data MaxInterface) (ok bool, max interface{}) {
	if data.Len() == 0 {
		return false, nil
	}
	if data.Len() == 1 {
		return true, data.Get(1)
	}
	max = data.Get(0)
	m := 0
	for i := 1; i < data.Len(); i++ {
		if data.Bigger(i, m) {
			max = data.Get(i)
			m = i
		}
	}
	return true, max
}

func main() {
	intslice := IntSlice{1, 2, 44, 6, 44, 222}
	floatslice := FloatSlice{1.99, 3.14, 24.8}
	group := PersonSlice{
		Person{name: "Jack", age: 24},
		Person{name: "Bob", age: 23},
		Person{name: "Bauer", age: 104},
		Person{name: "Paul", age: 44},
		Person{name: "Sam", age: 34},
		Person{name: "Lice", age: 54},
		Person{name: "Karl", age: 74},
		Person{name: "Lee", age: 4},
	}

	_, m := Max(intslice)
	fmt.Println("The biggest integer in islice is :", m)
	_, m = Max(floatslice)
	fmt.Println("The biggest float in fslice is :", m)
	_, m = Max(group)
	fmt.Println("The oldest person in the group is:", m)
}
