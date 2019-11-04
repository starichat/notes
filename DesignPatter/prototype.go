package main

import ("fmt")
// 原型模式代码实践

// 以简历克隆为例进行实现
type Workexpersience struct {
	content string
	workingyears int
}

type Resume struct {
	Workexpersience
	name string
	age int
	
}

func (w *Resume) SetWorkexp(content string, workingyears int) {
	w.content = content
	w.workingyears = workingyears
}

func (r *Resume) SetResume (name string, age int) {

	r.name = name 
	r.age = age
}

func (r *Resume) Clone() *Resume{
	res := *r
	return &res
	
}

func (r *Resume) Display() {
	fmt.Println(r.name, r.age)
	fmt.Println("工作经验：",r.content,r.workingyears)
}

func main() {
	r := new(Resume)
	r.SetWorkexp("安卓开发",1)
	r.SetResume("张三",12)

	r1:=r.Clone()
	r1.SetWorkexp("go 语言开发",2)

	r.Display()
	r1.Display()
}