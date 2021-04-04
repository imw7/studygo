package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件

// 利用file.Read读取文件
func readFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	// 记得关闭文件
	defer func() { _ = file.Close() }()
	// 读文件
	// var data = make([]byte, 128) // 指定读的长度
	var data [128]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("read from file failed, err:", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(data[:n]))
	}
}

// 利用bufio按行读取文件
func readByBufio(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer func() { _ = file.Close() }()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read line failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// 利用ioutil直接读取整个文件
func readByIoutil(name string) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(file))
}

func main() {
	readFile("./main.go")
	readByBufio("./main.go")
	readByIoutil("./main.go")
}
