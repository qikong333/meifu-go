package apis

import (
	_ "github.com/go-sql-driver/mysql"	// 配置mysql时需要单独引入
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/gin-gonic/gin"
)


var db *gorm.DB


func init() {
	var dbSrc = "root:123456@tcp(47.106.112.237:3306)/meifu?parseTime=true"
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

type conf struct {
	Path    string `yaml:"path"`
	Enabled bool   `yaml:"enabled"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}