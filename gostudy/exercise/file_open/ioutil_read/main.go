package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filepath := "/Users/haoyunpeng/gowork/src/github.com/gobible/gostudy/exercise/file_open/ttt.txt"
	file, err := ioutil.ReadFile(filepath)
	// if err == io.EOF {
	// 	fmt.Println("文件读取结束")
	// }
	if err != nil {
		fmt.Println("文件读取错误", err)
	}
	fmt.Println(string(file))

}
