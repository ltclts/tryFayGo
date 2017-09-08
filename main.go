package main

import (
	"app/router"
	"github.com/henrylee2cn/faygo"
)

func main() {
	router.Route(faygo.New("app"))
	faygo.Run()
}
