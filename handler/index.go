package handler

import (
	"github.com/henrylee2cn/faygo"
)

/*
Index
*/
var Index = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	return ctx.Render(200, faygo.JoinStatic("view/handler/index/index.html"), faygo.Map{
		"title":   "go",
		"content": "Welcome",
	})
})
