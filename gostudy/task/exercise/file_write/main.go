package main

import (
	"fmt"
	"os"
)

func main() {
	fpath := "/Users/haoyunpeng/gowork/src/github.com/gobible/gostudy/exercise/file_write/www.txt"
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("文件打开失败，失败原因：", err)
		return
	}
	defer f.Close()
	str := "hello file writer"
	n, err := f.Write([]byte(str))
	if err != nil {
		fmt.Println("文件写入失败，失败原因：", err)
		return
	}
	fmt.Println(n)

}
