// go 文件读写
// 1. os包实现复制
// 2. ioutil包实现剪贴
// 3. 基于bufio包实现复制和粘贴功能
// 4. 选取一个方式实现按列读取功能

一门语言的核心便是对文件进行操作读写，在本地机器上进行操作是对本地磁盘文件的读写，在web上文件读写也是一项重要的功能。

下面我们通过go提供的三种方式对文件进行读写，并分析其优缺点。

1. os包
```
fd, err := os.Open(filename) //打开文件，获取文件对象
buf := make([]byte, 1024)  // 创建一个1024字节的buf
n, err := fd.Read(buf)  // 读取文件内容
```
```
fo,err := os.Open(name,flag,perm)//创建文件并获取文件对象
io.WriteString(f, data) //写入字符串
```
2. ioutil
```
fd, err := os.Open(filename)
r := bufio.NewReader(fd)
buf :=make([]byte,1024)
n, err:= r.Read(buf)

```
```
ioutil.WriteFile(filename,bytes,perm) //写入字节数组
```
3. bufio
```
fd,err := os.Open(filename)
fd,err := ioutil.ReadAll(fd)
```
```
fo,err := os.Open(name,flag,perm)//创建文件并获取文件对象
w:=bufio.NewWriter(fo)
w.WriteString(data)
```

4. file
fo,err := os.Open(name,flag,perm)//创建文件并获取文件对象
fo.Write(bytes)