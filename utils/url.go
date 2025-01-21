package utils

import "strings"

type UrlBuilder struct {
	url string
}

func (ctx *UrlBuilder) Add(key, value string) {
	if strings.Contains(ctx.url, "?") {
		ctx.url = ctx.url + "&" + key + "=" + value
	} else {
		ctx.url = ctx.url + "?" + key + "=" + value
	}
}

func (ctx *UrlBuilder) ToString() string {
	return ctx.url
}

func CreatePath(url string) UrlBuilder {
	return UrlBuilder{url: url}
}
