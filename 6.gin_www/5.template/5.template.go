package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const dir = "/root/www/go_www/src/go_study/6.gin_www/5.template/"

//http://192.168.1.177:8080/index
func test1() {
	router := gin.Default()

	//加载模板
	router.LoadHTMLGlob(dir + "templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//定义路由
	router.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")

}

//http://192.168.1.177:8080/posts/index
//
func test2() {
	router := gin.Default()
	router.LoadHTMLGlob(dir + "templates/**/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.Run(":8080")
}

func main() {
	//test1()
	test2()
}
