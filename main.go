package main

import (
	"fleurui_api/client"
)

func main() {
	request := client.NewClient("6d7e700ebfa8a46b19f5b52076073b03", "t358eo9k9q9zn5uhhso2we0k0qtloozv")
	//request.ImageUpload(model.Image{FileBase64: "131231"})
	request.Dirs()
}
