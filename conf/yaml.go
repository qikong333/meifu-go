package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Serve    Serve
	Listen string   `yaml:"listen"`
	Cors string   `yaml:"cors"`

}

type Serve struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Sqlname string `yaml:"sqlname"`
}



func (c *Conf) GetConf() *Conf {

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