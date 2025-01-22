package utils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

var imageType = [3]string{"jpg", "jpeg", "png"}

func ToBase64(file *os.File) (string, string) {
	suffix, err := getFileSuffix(file.Name())
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}
	bufio.NewReader(file)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
	}(file)
	base64Encoding := base64.StdEncoding.EncodeToString(bytes)
	return base64Encoding, suffix
}

func getFileSuffix(fileName string) (string, error) {
	splits := strings.Split(fileName, ".")
	suffix := "." + splits[len(splits)-1]
	flag := false
	for _, s := range imageType {
		if s == suffix {
			flag = true
		}
	}
	if !flag {
		return "", fmt.Errorf("不支持的文件类型")
	}
	return suffix, nil
}
