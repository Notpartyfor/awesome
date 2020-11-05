package main

import (
	"gin/gw/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	server.InitService(r)
	r.Run(":8000")
}
