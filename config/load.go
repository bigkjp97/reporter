package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/**
@Author: KJP
@Date: 2021-09-18
获取并解析配置文件
*/

/**
将文件内容以字符串读出来
*/
func load(f *os.File) string {
	var j string

	p := bufio.NewReader(f)
	// 按行读取配置json，读完退出
	for {
		l, _, err := p.ReadLine()
		if err == io.EOF {
			break
		}
		j += string(l)
	}
	return j
}

func Load(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Printf("配置文件获取失败 [Error]: %s\n", err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("配置文件关闭异常 [Error]: %s\n", err.Error())
		}
	}(file)
	json := load(file)
	fmt.Println(json)
	return json
}
