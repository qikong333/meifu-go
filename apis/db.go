package apis

import (
	_ "github.com/go-sql-driver/mysql"	// 配置mysql时需要单独引入
	"github.com/jinzhu/gorm"
	"log"
	"github.com/gin-gonic/gin"
)


var db *gorm.DB


func init() {
	var dbSrc = "root:123456@tcp(127.0.0.1:3306)/test?parseTime=true"
	newDb, err := gorm.Open("mysql", dbSrc)
	if err != nil {
		log.Fatalln(err)
	}

	db = newDb

	sqlDb := db.DB()
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(15)

	if gin.IsDebugging(){
		db = db.Debug()
	}

	//defer 	db.Close()
}

func DB() *gorm.DB {
	return db.New()
}