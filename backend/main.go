package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"student-management-system/db"
	"student-management-system/handlers"
)

func main() {
	// 初始化 MongoDB 连接
	client := db.InitMongoDB()
	defer client.Disconnect(nil)

	// 创建 Gin 路由
	router := gin.Default()

	// 路由组
	api := router.Group("/api/v1")
	{
		api.GET("/students", handlers.GetStudentsHandler)
		api.POST("/students", handlers.AddStudentHandler)
	}

	// 启动服务器
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
