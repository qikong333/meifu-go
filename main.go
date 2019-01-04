package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"meifu/apis"
	"meifu/upload"
	"net/http"
	"strings"
)
 
func main() {

	r := gin.Default()
	// 	允许使用跨域请求  全局中间件
	r.Use(Cors())
	//	注册静态文件路径
	r.StaticFS(upload.GetImageFullPath(), http.Dir(upload.GetImageFullPath()))

	//	上传图片
	r.POST("/upLoad",apis.UploadImage)

	//	首页banner图操作
	r.GET("/bannerGet",apis.GetBanner)
	r.POST("/bannerAdd",apis.AddBanner)
	r.GET("/bannerDelete",apis.DeleteBanner)

	////	news操作
	//r.GET("/newGet",apis.GetNew)
	//r.POST("/newAdd",apis.AddNew)
	//r.GET("/newDelete",apis.DeleteNew)

	r.Run() // listen and serve on 0.0.0.0:8080


}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "http://localhost:4200")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}






