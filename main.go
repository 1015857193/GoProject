package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	config "GoProject/ConfigClass"
	test "GoProject/TestClass"
)

// func main() {
// 	fmt.Print("hello world\n")

// 	f, _ := os.Create("gin.testlog")
// 	gin.DefaultWriter = io.MultiWriter(f)

// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message":     "success",
// 			"descprition": "这是一个尝试",
// 			"vendor":      "1",
// 		})
// 	})
// 	r.Run()
// }
// 绑定为 JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {

	gin.ForceConsoleColor()

	test.Log()
	config.ConfigLog()
	fmt.Println("有意思")

	router := gin.Default()

	// JSON 绑定示例 ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	})

	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// 绑定HTML表单的示例 (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		//这个将通过 content-type 头去推断绑定器使用哪个依赖。
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.GET("/getJson", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message": "恭喜你", "id": "123"})
		fmt.Println("请求成功了")
	})

	router.Run(":8080")

}
