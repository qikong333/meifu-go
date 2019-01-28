package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetNew(c *gin.Context)  {

	id := c.Query("id")
	fmt.Printf("输出id %d\n",id)
	fmt.Println(id)
	var new []New
	if len(id) != 0 {
		if err :=db.Where("id = ?", id).First(&new).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
	}else{
		if err :=db.Find(&new).Error; err != nil {	 //  SELECT * FROM users;
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
	}

	c.JSON(200, gin.H{
		"code":"200",
		"msg":"查询成功",
		"data":new,
	})

}

func AddNew(c *gin.Context)  {


	var new New	// 不能为指针

	new.Id = c.PostForm("id")
	new.Title = c.PostForm("title")
	new.Author =c.PostForm("author")
	new.Content = c.PostForm("content")
	new.Source = c.PostForm("source")
	new.Time = time.Now()

	// 添加
	if len(new.Id) != 0 {
		if err := db.Model(&new).Where("id = ?", new.Id).Save(&new).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			c.JSON(404, gin.H{

				"msg":"添加失败",
				"data":new,
			})
		}
	}else{
		if err := db.Create(&new).Error; err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
			c.JSON(404, gin.H{

				"msg":"添加失败",
				"data":new,
			})
		}
	}

	 // 保存


	c.JSON(200, gin.H{
		"code":200,
		"msg":"添加成功",
		"data":new,
	})
}

func DeleteNew(c *gin.Context)  {
	id := c.Query("id")
	if err := db.Where("id = ?", id).Delete(&New{}).Error; err != nil {  // Banner{} =>结构体名称
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