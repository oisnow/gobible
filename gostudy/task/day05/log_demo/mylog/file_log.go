package mylog

import (
	"fmt"
	"os"
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
		return
	}
	f.Logfile = file
}

//Debug 记录Debug日志
func (f *FileLog) Debug(msg string) {
	//写日志
	// 需要丰富日志格式，时间、日志级别，哪一行、哪一个函数
	f.Logfile.WriteString(msg)
	return
}
