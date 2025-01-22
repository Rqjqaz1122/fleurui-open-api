package client

import (
	"encoding/json"
	"fmt"
	"github.com/Rqjqaz1122/fleurui-open-api/consts"
	"github.com/Rqjqaz1122/fleurui-open-api/http"
	"github.com/Rqjqaz1122/fleurui-open-api/model"
	"github.com/Rqjqaz1122/fleurui-open-api/utils"
	"os"
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

func (ctx *Client) ImageUpload(image model.Image) model.Result {
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
	resp := ctx.sendRequest(dataParam)
	return response(resp)
}

func (ctx *Client) FileUpload(file *os.File, dirId int, quality float64) model.Result {
	dataParam := DataParam{}
	dataParam.Method = consts.UploadImage
	dataParam.InterfaceName = consts.Image
	fileBase64, suffix := utils.ToBase64(file)
	dataParam.ExtData = extData{
		DirId:      dirId,
		FileBase64: fileBase64,
		Suffix:     suffix,
		Quality:    quality,
	}
	resp := ctx.sendRequest(dataParam)
	return response(resp)
}

func (ctx *Client) Images() model.Result {
	dataParam := DataParam{}
	dataParam.Method = consts.ListImage
	dataParam.InterfaceName = consts.Image
	resp := ctx.sendRequest(dataParam)
	return response(resp)
}

func (ctx *Client) Random(dirId int) model.Result {
	dataParam := DataParam{}
	dataParam.Method = consts.RandomImage
	dataParam.InterfaceName = consts.Image
	if dirId == 0 {
		dataParam.ExtData = extData{DirId: -1}
	}
	resp := ctx.sendRequest(dataParam)
	return response(resp)
}

func (ctx *Client) CreateDir(parentDirId int, dirName string) model.Result {
	dataParam := DataParam{}
	dataParam.InterfaceName = consts.Image
	dataParam.Method = consts.CreateDir
	dataParam.ExtData = extData{
		ParentDirId: parentDirId,
		DirName:     dirName,
	}
	resp := ctx.sendRequest(dataParam)
	return response(resp)
}

func (ctx *Client) Dirs() model.Result {
	param := DataParam{}
	param.InterfaceName = consts.Image
	param.Method = consts.ListDir
	resp := ctx.sendRequest(param)
	return response(resp)
}

func (ctx *Client) Bucket() model.Result {
	param := DataParam{}
	param.InterfaceName = consts.User
	param.Method = consts.GetBucket
	resp := ctx.sendRequest(param)
	return response(resp)
}

func (ctx *Client) CreateBucket(name string) model.Result {
	param := DataParam{}
	param.InterfaceName = consts.User
	param.Method = consts.CreateBucket
	param.ExtData.BucketName = name
	resp := ctx.sendRequest(param)
	return response(resp)
}

func (ctx *Client) UserInfo() model.Result {
	param := DataParam{}
	param.InterfaceName = consts.User
	param.Method = consts.UserInfo
	resp := ctx.sendRequest(param)
	return response(resp)
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

func response(resp string) model.Result {
	result := model.Result{}
	unErr := json.Unmarshal([]byte(resp), &result)
	if unErr != nil {
		fmt.Println("解析json失败", unErr)
	}
	return result
}
