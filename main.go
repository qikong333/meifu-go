package main

import (
	"meifu/apis"
	"github.com/gin-gonic/gin"
	"net/http"
	"meifu/upload"
)
 
func main() {


	r := gin.Default()
	//注册静态文件路径
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
 