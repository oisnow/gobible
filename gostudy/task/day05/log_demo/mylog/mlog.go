package mylog

const (
	//DEBUG 定义debug级别的函数level数值为0
	DEBUG = iota
	//TRACE 定义debug级别的函数level数值为1
	TRACE
	//INFO 定义debug级别的函数level数值为2
	INFO
	//ERROR 定义debug级别的函数level数值为3
	ERROR
)

// GetLogLevelStr 获取不同level数值对应的日志级别
func GetLogLevelStr(level int) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "ERROR"
	default:
		return "DEBUG"
	}

}
