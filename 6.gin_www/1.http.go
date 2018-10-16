package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func test1() {
	router := gin.Default()
	//直接输出123
	router.GET("/", func(c *gin.Context) {
		c.String(200, `123`)
	})

	//带http 200头的,http体是123
	router.GET("/http", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `123`)
	})

	http.ListenAndServe(":8080", router)
}

func test2() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

//curl 127.0.0.1:8080
//hello /

//curl 127.0.0.1:8080/string/ybx
//hello ybx

func test3() {
	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	router := gin.Default()
	//创建不带中间件的路由：
	//r := gin.New()

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Hello /")
	})
	router.GET("/string/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println("Hello ", name)
	})

	//	http.ListenAndServe(":8080", router)
	router.Run(":8080")
}

// curl http://localhost:8080/welcome
//Hello name:Guest lastname:

// curl "http://localhost:8080/welcome?name=1111&lastname=2222"
//Hello name:ningskyer lastname:
func test4() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
		//lastname := c.DefaultQuery("lastname","Guest2")// 是 c.Request.URL.Query().Get("lastname") 的简写
		fmt.Printf("Hello name:%s lastname:%s\n", name, lastname)
	})
	router.Run(":8080")
}

//curl -d "type=1111&msg=2222&title=3333" "127.0.0.1:8080/form"
func test5_form() {
	router := gin.Default()

	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert") //可设置默认值
		msg := c.PostForm("msg")
		title := c.PostForm("title")
		fmt.Printf("type is %s, msg is %s, title is %s", type1, msg, title)
	})
	router.Run(":8080")
}

//http://192.168.1.177:8080/someGroup/someGet
func test6() {
	router := gin.Default()
	someGroup := router.Group("/someGroup")
	{
		someGroup.GET("/someGet", func(c *gin.Context) {
			fmt.Println("1111")
		})
	}
	router.Run(":8080")
}

func main() {
	test1()
	//test2()
	//test3()
	//test4()
	//test5_form()
	//test6()
}
