package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context)  {

	var product []Product	// 变量名的定义要和数据表名相同

	if err :=db.Find(&product).Error; err != nil {	 //  SELECT * FROM users;
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"code":"200",
		"msg":"查询成功",
		"data":product,
	})

}

func AddProduc(c *gin.Context)  {
	var product Product	// 不能为指针

	product.Content = c.PostForm("content")
	product.Name = c.PostForm("name")
	product.Img = c.PostForm("img")
	product.Info = c.PostForm("info")


	if err := db.Create(&product).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
		c.JSON(404, gin.H{
			"msg":"添加失败",
			"data":product,
		})
	}

	c.JSON(200, gin.H{
		"msg":"添加成功",
		"data":product,
	})
}

func DeleteProduct(c *gin.Context)  {
	id := c.Query("id")
	if err := db.Where("id = ?", id).Delete(&Product{}).Error; err != nil {  // Banner{} =>结构体名称
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