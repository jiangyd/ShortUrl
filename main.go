package main

import (
	"github.com/gin-gonic/gin"
	"shortUrl/control"
	"shortUrl/config"

)

func main()  {
	r:=gin.Default()

	r.LoadHTMLGlob("templates/*")
	//r.GET("/",control.Index)
	r.POST("/short",control.Short)
	r.GET("/*path",control.Redirct)
	r.Run(config.Cfg.Bind)
}



