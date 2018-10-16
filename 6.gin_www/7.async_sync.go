package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//http://192.168.1.177:8080/long_async
//http://192.168.1.177:8080/long_sync
func main() {
	r := gin.Default()
	//1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 注意使用只读上下文
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	//2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		// 注意可以使用原始上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
