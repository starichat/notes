// go 文件读写
// 1. os包实现复制
// 2. ioutil包实现剪贴
// 3. 基于bufio包实现复制和粘贴功能
// 4. 选取一个方式实现按列读取功能

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func RawRead(filename string) string {
	start := time.Now()
	fd, err := os.Open(filename) // 获取文件对象
	if err != nil {
		panic(err)
	}
	defer func() {
		fd.Close()
		fmt.Printf("[RawRead] cost time %v \n", time.Now().Sub(start))
	}()
	var data []byte
	buf := make([]byte, 1024)
	//fmt.Println(buf)
	for {
		n, err := fd.Read(buf) //读取文件内容到buf上去
		//	fmt.Println(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		data = append(data, buf[:n]...)
		if n == 0 {
			break
		}
	}
	return string(data)
}

func IOutilRead(filename string) string {
	start := time.Now()
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		fd.Close()
		fmt.Printf("[IOutilRead] cost time %v \n", time.Now().Sub(start))
	}()
	data, err := ioutil.ReadAll(fd)
	return string(data)
}
func bufferRead(filename string) []byte {
	start := time.Now()
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		fd.Close()
		fmt.Printf("[BufferRead] cost time %v \n", time.Now().Sub(start))
	}()
	r := bufio.NewReader(fd)
	var data []byte
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)

	}
	return data
}

func RawWrite(desc string, data string) {
	start := time.Now()
	fo, err := os.Create(desc) //创建文件并获取文件对象
	if err != nil {
		panic(err)
	}
	defer func() {
		fo.Close()
		fmt.Printf("[RawWrite] cost time %v \n", time.Now().Sub(start))
	}()
	n, err := io.WriteString(fo, data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入了 %d 个字节", n)
}
func IoutilWrite(desc string, bytes []byte) {
	start := time.Now()
	err := ioutil.WriteFile(desc, bytes, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Printf("[IOWrite] cost time %v \n", time.Now().Sub(start))
	}()

}

func BufferWrite(desc string, bytes string) {
	start := time.Now()
	fo, err := os.Create(desc) //创建文件并获取文件对象
	if err != nil {
		panic(err)
	}
	defer func() {
		fo.Close()
		fmt.Printf("[BufferWrite] cost time %v \n", time.Now().Sub(start))
	}()
	w := bufio.NewWriter(fo)
	n, err := w.WriteString(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入了%d个字节", n)
	w.WriteByte('B')
	fmt.Printf("写入了%d个字节", n)
	w.Flush()

}

func FileWriter(desc string, bytes []byte) {
	start := time.Now()
	fo, err := os.Create(desc) //创建文件并获取文件对象
	if err != nil {
		panic(err)
	}
	defer func() {
		fo.Close()
		fmt.Printf("[FileWrite] cost time %v \n", time.Now().Sub(start))
	}()
	n, err := fo.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入%d个字节", n)
}

func FileWritestring(desc string, data string) {
	start := time.Now()
	fo, err := os.Create(desc) //创建文件并获取文件对象
	if err != nil {
		panic(err)
	}
	defer func() {
		fo.Close()
		fmt.Printf("[FileWritestring] cost time %v \n", time.Now().Sub(start))
	}()
	n, err := fo.WriteString(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("写入%d个字节", n)
}

func main() {
	RawRead("test.txt")
	IOutilRead("test.txt")
	bufferRead("test.txt")
	RawWrite("t2.txt", RawRead("test.txt"))
	IoutilWrite("t3.txt", bufferRead("test.txt"))
	BufferWrite("t4.txt", RawRead("test.txt"))
	FileWriter("t5.txt", bufferRead("test.txt"))
	FileWritestring("t6.txt", RawRead("test.txt"))
}
