package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter() (r *gin.Engine) {
	r = gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		fmt.Println("running...")
		fmt.Println(err)
		fmt.Println("running...")
	}))

	r.GET("/ping", func(c *gin.Context) {
		ResponseJson(c, PING, nil)
	})
	r.GET("/", func(c *gin.Context) {
		panic("error")
		//ResponseJson(c, ERR, nil)
	})

	r.Group("")

	r.NoRoute(func(c *gin.Context) {
		ResponseJson(c, ERR, nil)
	})
	r.NoMethod(func(c *gin.Context) {
		ResponseJson(c, ERR, nil)
	})
	//r.

	return r
}

var emptyData = struct{}{}

func ResponseJson(c *gin.Context, stat HttpStat, data interface{}) {
	if data == nil {
		data = &emptyData
	}
	c.JSON(200, gin.H{
		"code": stat.code,
		"msg":  stat.msg,
		"data": data,
	})
}
