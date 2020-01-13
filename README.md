# GoBible
### exercise of https://www. gitbook.com/book/yar999 /gopl-zh/details 
### go语言学习之go语言圣经的练习题

### 测试类

因为一个包中只能有一个人func main，故只有helloword.go 中使用了正常的go文件编写，其他都使用测试类进行练习
```go
package main

import "testing"

func TestEcho(t *testing.T) {
	t.Log("my is test go program")
}
```

