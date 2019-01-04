package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 配置mysql时需要单独引入
	"meifu/pkg/e"
	"net/http"
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
	var banner Banner

	banner.Name  = c.PostForm("name")
	banner.Name  = c.PostForm("image")

	if err := db.Create(&banner).Error; err != nil {
		code = e.ERROR
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(code, gin.H{
			"msg":"添加失败",
			"data":banner,
		})
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