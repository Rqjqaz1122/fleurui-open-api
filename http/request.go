package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Client *http.Client
}

// NewHttpClient 初始化http客户端
func NewHttpClient() *Request {
	return &Request{
		Client: &http.Client{},
	}
}

func (ctx *Request) Get(url string) (string, error) {
	resp, err := ctx.Client.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭请求失败")
		}
	}(resp.Body)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (ctx *Request) Post(url string, contentType string, data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	resp, err := ctx.Client.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("关闭请求失败")
		}
	}(resp.Body)
	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(byteData), nil
}
