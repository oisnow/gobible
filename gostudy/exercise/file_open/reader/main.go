package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "/Users/haoyunpeng/gowork/src/github.com/gobible/gostudy/exercise/file_open/ttt.txt"
	f, err := os.Open(filepath)

	if err != nil {
		fmt.Println("文件不存在", err)
		return
	}
	defer f.Close()
	filetmp := make([]byte, 128, 128)

	for {
		line, err := f.Read(filetmp)
		if err == io.EOF {
			fmt.Println("文件已经读完")
			return
		}
		if err != nil {
			fmt.Println("打开文件错误", err)
			return
		}
		fmt.Printf("读取了%d个字节 \n", line)
		// fmt.Println(line)
		fmt.Println(string(filetmp))
	}
}
