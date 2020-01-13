package main

import (
	"os"
	"strings"
	"testing"
)

//可以使用strings.join进行字符拼接工作
func TestEcho3(t *testing.T) {
	t.Log(os.Args[1:])
	t.Log(strings.Join(os.Args[1:], "@@"))
}
