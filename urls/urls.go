package urls

import (
	"github.com/TobiasYin/dcc/handlers"
	"github.com/gin-gonic/gin"
)

func InitUrls(e *gin.Engine){
	// Register Test Url
	e.GET("/ping", handlers.Ping)
}
