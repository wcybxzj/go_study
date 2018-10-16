package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//curl -X POST -H "Content-Type:application/json" -d "{\"user\": \"manu\",\"password\":\"123\"}" http://192.168.1.177:8080/loginJSON
//curl -d "user=manu&password=123" "http://192.168.1.177:8080/loginForm"
//curl -d "user=manu&password=123" "http://192.168.1.177:8080/login"
func main() {
	router := gin.Default()
	// 绑定JSON的例子 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login

		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// 绑定普通表单的例子 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 根据请求头中 content-type 自动推断.
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// 绑定多媒体表单的例子 (user=manu&password=123)
	router.POST("/login", func(c *gin.Context) {
		var form Login
		// 你可以显式声明来绑定多媒体表单：
		// c.BindWith(&form, binding.Form)
		// 或者使用自动推断:
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
