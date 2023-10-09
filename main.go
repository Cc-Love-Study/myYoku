package main

import (
	"fmt"
	"myYoku/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello gin")
	// 声明一个路由
	r := gin.Default()

	userController := controllers.NewUserController(r, "user")
	userController.InitUserController()

	r.Run(":8080")
}
