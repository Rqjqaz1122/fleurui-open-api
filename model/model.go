package model

type Image struct {
	FileBase64 string
	Suffix     string
	ImageName  string
	DirId      int
	Quality    float64
}

type Result struct {
	code int
	msg  string
	data string
}
