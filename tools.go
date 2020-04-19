// Package tools 一些常用的功能函数
package tools

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// 根据系统信息判断生产环境还是开发环境
func IsDev() (isDev bool) {
	switch runtime.GOOS {
	case "windows", "darwin":
		isDev = true
	default:
		isDev = false
	}
	return
}

// 获取当前程序的绝对路径
func RealPath() (dir string) {
	if IsDev() {
		dir, _ = os.Getwd()
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	dir = dir + string(os.PathSeparator)
	return
}

// 字符串数组中是否包含某个值
func ContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// 把json文件中的内容解析到指定type
func JsonToStruct(path string, s interface{}) {
	f, err := os.Open(path)
	ErrCheck(err)
	defer f.Close()

	err = json.NewDecoder(f).Decode(s)
	ErrCheck(err)
	return
}

// 错误处理
func ErrCheck(err error) {
	if err != nil {
		log.Println(err)
	}
}

// []unit8转为string
func B2S(bs []uint8) string {
	var ba []byte
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
