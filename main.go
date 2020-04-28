package main

import (
	"github.com/TobiasYin/dcc/middleware"
	// Init DataBase Model
	_ "github.com/TobiasYin/dcc/models"
	"github.com/TobiasYin/dcc/urls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Notice: Make Sure Init MiddleWare before Urls
	middleware.InitMiddleWare(r)
	urls.InitUrls(r)
	if e := r.Run(); e != nil {
		panic(e)
	}
}
