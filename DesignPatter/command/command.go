package main

// 只有一个功能:执行
type Command interface {
	execute()
}

// receiver
type Receiver interface {
	receive()
}

// 各个子类重写执行的方法

//

func main() {

	// 接受命令

	// 执行命令

}
