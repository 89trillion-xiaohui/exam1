package main


import (
	"./mesg"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

	//"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"

)



var jsonPath string	//存储json路径

func main() {
	//gin
	r := gin.Default()
	r.GET("/config.army.model.json", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		_, err := c.Writer.Write( []byte("???"))
		if err != nil {
			return
		}
	})
	//读取ini文件
	cfg, err := ini.Load("/Users/wenxiaohui/Downloads/1. 解析配置文件/app.ini")
	if err != nil {
		fmt.Printf("Failed to read file : %v", err)
	 	os.Exit(1)
	}
	//获取客户端端口
	httpPort := cfg.Section("server").Key("HttpPort").Value()
	fmt.Println(httpPort)

	//json路径
	jsonPath = "/Users/wenxiaohui/Desktop/Go/Hello/exam/config.army.model.json"

	//读取json文件
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("?")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(string(byteValue))

	//解析json
	var s = mesg.Soldier{}
	err = json.Unmarshal([]byte(byteValue), &s)
	if err != nil {
		return
	}
	fmt.Println(s)



	r.Run(httpPort)
}
