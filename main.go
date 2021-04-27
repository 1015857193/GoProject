package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "GoProject/docs"
	//_ "github.com/swaggo/gin-swagger/example/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

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

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath /api/v1/
func main() {

	//main
	TestMain()
	//关机重启
	//testShutdown()

}

// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce json
// @Param   {"message": "恭喜你", "id": "123"}
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /api/v1/getJson
func TestMain() {

	gin.ForceConsoleColor()

	test.Log()
	config.ConfigLog()
	fmt.Println("有意思")

	router := gin.Default()

	// JSON 绑定示例 ({"user": "manu", "password": "123"})
	router.POST("/api/v1/loginJSON", func(c *gin.Context) {
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
	//xml 绑定示例
	router.POST("/api/v1/loginXML", func(c *gin.Context) {
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
	router.POST("/api/v1/loginForm", func(c *gin.Context) {
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

	//get方法
	router.GET("/api/v1/getJson", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message": "恭喜你", "id": "123"})
		fmt.Println("请求成功了")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func testShutdown() {

	router := gin.Default()
	router.GET("/shutdown", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
