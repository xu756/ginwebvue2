package config

import (
	"fmt"
	"os"
)

// CreateDir 创建文件夹
func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) { // 判断文件夹是否存在
		err := os.MkdirAll(path, 0755) // 创建文件夹
		fmt.Println("创建文件夹", path, "成功")
		if err != nil {
			return // 创建失败
		}
	}
}
