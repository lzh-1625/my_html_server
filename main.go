package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("/public","./public")
	r.GET("/",func(ctx *gin.Context) {
		ctx.HTML(200,"/index.html", nil)
	})
	r.Run(":8585")
}