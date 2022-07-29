/**
*  @Author: xu756 (https://github.com/xu756)
*  @Email: 756334744@qq.com
*  @Date: 2022/7/29 22:36
*  @To:		读取数据
 */

package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	// Redis 配置
	Redis redisConfig `json:"redis"`
	// 发送邮件配置
	Email emailConfig `json:"email"`
}

var InitData Data

func init() {
	file, err := os.Open("./config.json")
	if err != nil {
		fmt.Printf("文件打开失败 [Err:%s]\n", err.Error())
		return
	}
	defer file.Close()
	// 读取文件内容
	var data Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		InitData = data
	}

}
