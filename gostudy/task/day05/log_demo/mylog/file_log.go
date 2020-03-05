package mylog

import (
	"fmt"
	"os"
	"time"
)

// FileLog 文件读写日志类型的结构体
type FileLog struct {
	LogFileName string
	LogFilePath string
	LogLevel    int
	Logfile     *os.File
}

// NewFileLog 实例化文件读写类型日志的示例,一个构造函数
// LogLevel：debug(1)、trace(2)、info(3)、error(3)
func NewFileLog(LogLevel int, LogFileName, LogFilePath string) *FileLog {
	Fileobj := &FileLog{
		LogFileName: LogFileName,
		LogFilePath: LogFilePath,
		LogLevel:    LogLevel,
	}

	Fileobj.initLogpath()
	return Fileobj
}

//初始化日志文件句柄
func (f *FileLog) initLogpath() {
	filepath := fmt.Sprintf("%s%s", f.LogFilePath, f.LogFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		// fmt.Println("打开文件错误，错误原因：", err)
		panic(fmt.Sprintf("open %s file error", filepath))
	}
	f.Logfile = file
}

// GetTime 获取当前时间并格式化为 2006/01/02 15:04:05.000
func GetTime() string {
	t := time.Now()
	timestr := t.Format("2006/01/02 15:04:06.000")
	return timestr
}

//Debug 记录Debug日志
func (f *FileLog) Debug(format string, arg ...interface{}) {

	if f.LogLevel > 0 {
		return
	}
	//写日志
	// 需要丰富日志格式，时间、日志级别，哪一行、哪一个函数
	files, funcnames, line := GetCallInfo()
	fmt.Println("")
	timenow := GetTime()
	format = fmt.Sprintf("[时间：%s ,Level：%s 文件名：%s ，方法名：%s ，行号：%d] ,日志内容：%s", timenow, GetLogLevelStr(f.LogLevel), files, funcnames, line, format)
	fmt.Fprintf(f.Logfile, format, arg...)
	fmt.Fprintln(f.Logfile)
	return
}
