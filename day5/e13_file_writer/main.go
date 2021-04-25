package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 打开文件写内容

func writeFile(name string) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func() { _ = file.Close() }()
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read input failed, err:", err)
			return
		}
		msg = strings.Trim(msg, "\r\n")
		if strings.ToUpper(msg) == "EXIT" {
			break
		}
		// write
		_, _ = file.Write([]byte("Write: " + msg + "\n"))
		// writeString
		_, _ = file.WriteString("WriteString: " + msg + "\n")
	}
}

func writeWithBufio(name string) {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer func() { _ = file.Close() }()
	reader := bufio.NewReader(os.Stdin)
	// 创建一个写的对象
	writer := bufio.NewWriter(file)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read input failed, err:", err)
			return
		}
		msg = strings.Trim(msg, "\r\n")
		if strings.ToUpper(msg) == "EXIT" {
			break
		}
		_, _ = writer.WriteString(msg + "\n") // 写到缓存中
		_ = writer.Flush()                    // 将缓存中的内容写到文件中
	}
}

func writeWithIoutil(name string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read input failed, err:", err)
			continue
		}
		data = strings.Trim(data, "\r\n")
		if strings.ToUpper(data) == "EXIT" {
			break
		}
		err = ioutil.WriteFile(name, []byte(data), 0777)
		if err != nil {
			fmt.Println("write file failed, err:", err)
			return
		}
	}
}

func main() {
	writeFile("xx1.txt")
	writeWithBufio("xx2.txt")
	writeWithIoutil("xx3.txt")
}
