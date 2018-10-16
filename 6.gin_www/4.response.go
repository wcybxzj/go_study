package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user" xml:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		c.JSON(http.StatusOK, msg)

		//c.XML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		//c.XML(http.StatusOK, msg)

		//c.YAML(http.StatusOK, gin.H{"user": "Lena", "Message": "hey", "Number": 123})
		//c.YAML(http.StatusOK, msg)

	})

	router.Run(":8080")
}
