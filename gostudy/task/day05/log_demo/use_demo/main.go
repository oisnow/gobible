package main

import (
	"github.com/gobible/gostudy/task/day05/log_demo/mylog"
)

func main() {
	var log = mylog.NewFileLog(1, "debug.log", "gostudy/task/day05/log_demo/use_demo/log")
	log.Debug("I'm a Debug Log")
}
