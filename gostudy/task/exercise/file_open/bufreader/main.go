package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "/Users/haoyunpeng/gowork/src/github.com/gobible/gostudy/exercise/file_op/ttt.txt"

	//文件读取
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("打开文件错误，错误内容：", err)
		return
	}
	defer file.Close()

	//读取内容
	filereader := bufio.NewReader(file)
	count := 1
	for {

		str, err := filereader.ReadString('.')
		if err == io.EOF {
			fmt.Println("\n--------------文件已读取完毕--------------")
			fmt.Printf("文件内容共%d行", count)
			return
		}
		if err != nil {
			fmt.Println("读取内容出错,错误内容：", err)
			return
		}
		fmt.Print(str)
		count++
	}

}
