package client

import (
	"fleurui_api/consts"
	"fleurui_api/http"
	"fleurui_api/model"
	"fleurui_api/utils"
	"fmt"
	"time"
)

const url = "https://www.api.wrqj.top/api/v1/openApi"

var client = http.NewHttpClient()

type Client struct {
	url string
}

type DataParam struct {
	Time          string  `json:"time"`
	InterfaceName string  `json:"interfaceName"`
	Method        string  `json:"method"`
	ExtData       extData `json:"extData"`
}

type extData struct {
	FileBase64  string  `json:"fileBase64"`
	DirId       int     `json:"dirId"`
	Suffix      string  `json:"suffix"`
	FileName    string  `json:"fileName"`
	Quality     float64 `json:"quality"`
	BucketName  string  `json:"bucketName"`
	ParentDirId int     `json:"parentDirId"`
	DirName     string  `json:"dirName"`
}

func NewClient(ak, sk string) Client {
	builder := utils.CreatePath(url)
	builder.Add("accessKey", ak)
	builder.Add("secretKey", sk)
	finalUrl := builder.ToString()
	return Client{
		url: finalUrl,
	}
}

func (ctx *Client) ImageUpload(image model.Image) {
	dataParam := DataParam{}
	dataParam.Method = consts.UploadImage
	dataParam.InterfaceName = consts.Image
	dataParam.ExtData = extData{
		FileBase64: image.FileBase64,
		Quality:    image.Quality,
		Suffix:     image.Suffix,
		DirId:      image.DirId,
		FileName:   image.ImageName,
	}
	fmt.Println("请求参数", *ctx)
	resp := ctx.sendRequest(dataParam)
	fmt.Println(resp)
}

func (ctx *Client) Images() {
	dataParam := DataParam{}
	dataParam.Method = consts.ListImage
	dataParam.InterfaceName = consts.Image
	resp := ctx.sendRequest(dataParam)
	fmt.Println(resp)
}

func (ctx *Client) Random(dirId int) {
	dataParam := DataParam{}
	dataParam.Method = consts.RandomImage
	dataParam.InterfaceName = consts.Image
	if dirId == 0 {
		dataParam.ExtData = extData{DirId: -1}
	}
	resp := ctx.sendRequest(dataParam)
	fmt.Println(resp)
}

func (ctx *Client) CreateDir(parentDirId int, dirName string) {
	dataParam := DataParam{}
	dataParam.InterfaceName = consts.Image
	dataParam.Method = consts.CreateDir
	dataParam.ExtData = extData{
		ParentDirId: parentDirId,
		DirName:     dirName,
	}
	ctx.sendRequest(dataParam)
}

func (ctx *Client) Dirs() {
	param := DataParam{}
	param.InterfaceName = consts.Image
	param.Method = consts.ListDir
	ctx.sendRequest(param)
}

func (ctx *Client) sendRequest(param DataParam) string {
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	param.Time = format
	resp, err := client.Post(ctx.url, "application/json", param)
	if err != nil {
		fmt.Println("fleurui_api调用失败")
	}
	return resp
}
