package mylog

import (
	"path"
	"runtime"
)

// GetCallInfo 获取函数调用的相关信息
func GetCallInfo() (files, funcnames string, line int) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	funcname := runtime.FuncForPC(pc).Name()
	funcnames = path.Base(funcname)
	files = path.Base(file)
	//打印调试
	// fmt.Println(funcname)
	// fmt.Println(file, line, ok)
	return
}
