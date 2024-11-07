package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	//"net/http"
	//"html/template"

	"myproject/post"
)

// 定义一个全局对象db
var db *sql.DB

func main() {

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	//跨域问题CORS设置
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access - Control - Allow - Origin", "*")
		c.Writer.Header().Set("Access - Control - Allow - Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access - Control - Allow - Headers", "Content - Type")
		c.Next()
	})

	r.LoadHTMLGlob("statics/html/*")

	r.Static("/css", "/root/myblog2/statics/css")
	r.Static("/images", "/root/myblog2/statics/images")
	r.Static("/js", "/root/myblog2/statics/js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/project", func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		if len(queryParams) == 0 {
			fmt.Println("没有获取到GET数据")
		} else {
			// 遍历所有查询参数并打印
			for k, v := range queryParams {
				fmt.Printf("获得GET数据的参数名: %s, 参数值: %v\n", k, v)
			}
		}
		c.HTML(200, "project.html", nil)
	})

	r.POST("/project", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println("获取POST数据失败:", err)
		} else {
			fmt.Printf("接收到的POST数据: %v\n", string(data))
			//打印数据
			//启动接收函数
			post.PostFunction(data)
		}
		c.HTML(200, "project.html", nil)
	})

	r.Run(":80") // 监听域名默认的80端口
}
