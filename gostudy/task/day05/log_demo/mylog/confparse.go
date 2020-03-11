package mylog

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

//Config 是一个日志的配置项目-结构体
type Config struct {
	Filepath string `conf:"file_path"`
	Filename string `conf:"file_name"`
	Loglevel string `conf:"log_level"`
	Maxsize  int64  `conf:"max_size"`
}

//ParseConfig 解析日志文件-函数
func ParseConfig(filename string, res interface{}) (err error) {

	//前提条件 res必须是一个ptr
	t := reflect.TypeOf(res)
	v := reflect.ValueOf(res).Elem()
	if t.Kind() != reflect.Ptr {
		err = errors.New("res必须是一个指针类型")
		return
	}
	//前提条件 res必须是一个ptr 并且是结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("res必须是一一个结构体指针类型")
		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("打开配置文件%s失败", filename)
		return err
	}
	//将文件按行分割成切片
	lineslice := strings.Split(string(data), "\r\n")
	for index, line := range lineslice {
		//去除空格
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		//判断:位置 判断是否符合规范
		equaIndex := strings.Index(line, ":")
		// ==-1 没找到
		if equaIndex == -1 {
			err = fmt.Errorf("配置%d行文件语法错误", index+1)
			return
		}
		//分割配置项目 获取key 与value
		key := line[:equaIndex]
		val := line[equaIndex+1:]
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		key = strings.Trim(key, "\"")
		val = strings.Trim(val, "\"")

		// fmt.Println(key, "+", val)
		if len(key) == 0 || len(val) == 0 {
			err = fmt.Errorf("配置%d行文件语法错误", index+1)
			return
		}
		//利用反射给res赋值
		//遍历结构体的与key比较  匹配成功进行赋值。
		for i := 0; i < v.NumField(); i++ {
			filed := v.Type().Field(i)
			tag := filed.Tag
			name := tag.Get("conf")
			// fmt.Println("name", name)
			// fmt.Println("key", key)
			if name == key {
				// fmt.Println(key, name, name == key)
				// fmt.Println(v.Field(i).Kind())
				filedval := v.FieldByName(filed.Name)
				switch v.Field(i).Kind() {
				case reflect.String:
					// fmt.Println("filedval", filedval)
					filedval.SetString(val)

				case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
					val64, _ := strconv.ParseInt(val, 10, 64)
					filedval.SetInt(val64)
				}

			}
			// fmt.Println(key, name, name == key)
		}
	}
	return
}
