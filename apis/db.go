package apis

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 配置mysql时需要单独引入
	"github.com/jinzhu/gorm"
	"log"
	"meifu/conf"
)


var db *gorm.DB


func init() {
	var c conf.Conf
	c.GetConf()
	var dbSrc =c.Serve.Username+ ":"+c.Serve.Password+"@tcp("+c.Serve.Addr+":"+c.Serve.Port+")/"+c.Serve.Sqlname+"?parseTime=true"
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
