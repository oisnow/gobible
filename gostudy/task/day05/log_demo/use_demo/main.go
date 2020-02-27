package main

import (
	"fmt"

	"github.com/gobible/gostudy/task/day05/log_demo/mylog"
)

func main() {
	var log = mylog.NewFileLog(mylog.DEBUG, "debug.log", "gostudy/task/day05/log_demo/use_demo/log")
	fmt.Println("输入123456的debug日志")
	log.Debug("输入日志为%d", 123456)
}
