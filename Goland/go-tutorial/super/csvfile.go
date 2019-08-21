package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readCsv(src string) []string {
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var result []string
	var s1, s2, s3, s4, s5 []string
	for {
		row, err := reader.ReadString('\n') //获取每一行数据
		//row = strings.Trimspace(row)        // 去除多余的空格
		if err == io.EOF {
			break
		}
		// 读取除第一行的数据
		result = append(result, row)
		// 将每一行的数据读取到不同的列
		s := strings.Split(row, ",")
		// 将，每一行的数据提取出来
		s1 = append(s1, s[0])
		s2 = append(s2, s[1])
		s3 = append(s3, s[2])
		s4 = append(s4, s[3])
		s5 = append(s5, s[4])

	}

	fmt.Println(s1, s2)

	return result

}

// func writeCsv(desc string) {
// 	file, _ := os.OpenFile(desc, os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	defer file.Close()
// 	w := csv.NewWriter(file)
// 	w.Write([]string{"id", "name", "company", "city", "salary"})
// 	w.Write([]string{"123", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"124", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"125", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"123", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"123", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"123", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"123", "234234", "345345", "234234", "adada"})
// 	w.Write([]string{"123", "234234", "345345", "234234", "adada"})
// 	w.Flush()
// }
func main() {
	//	writeCsv("userinfo.csv")
	fmt.Println(readCsv("userinfo.csv"))
}
