package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"	// 配置mysql时需要单独引入
	"meifu/pkg/e"
	"net/http"
	"meifu/upload"
)

func GetBanner(c *gin.Context)  {

	var banner []Banner	// 变量名的定义要和数据表名相同

	if err :=db.Find(&banner).Error; err != nil {	 //  SELECT * FROM users;
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.IndentedJSON(200, gin.H{
		"code":"200",
		"msg":"查询成功",
		"data":banner,
	})

}

func AddBanner(c *gin.Context)  {
	code := e.SUCCESS
	//data := make(map[string]string)
	var banner Banner

	banner.Name  = c.PostForm("name")
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		//logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": banner,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		//savePath := upload.GetImagePath()

		src := fullPath + imageName
		fmt.Println("图片名称："+imageName)
		fmt.Println("路径："+src)
		if ! upload.CheckImageExt(imageName) || ! upload.CheckImageSize(file) { // 检查后缀和大小
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)	//	检查图片
			if err != nil {
				//logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				fmt.Println(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				banner.Image = upload.GetImageFullUrl(imageName)
				//data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": banner,
	})
}


func DeleteBanner(c *gin.Context)  {
	id := c.Query("id")
	if err := db.Where("id = ?", id).Delete(&Banner{}).Error; err != nil {  // Banner{} =>结构体名称
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(404, gin.H{
			"msg":"删除失败",
		})
	}

	c.JSON(200, gin.H{
		"msg":"删除成功",
	})

}