package model

type Image struct {
	FileBase64 string
	Suffix     string
	ImageName  string
	DirId      int
	Quality    float64
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
