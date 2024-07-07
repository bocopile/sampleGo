package main

import (
	"github.com/gin-gonic/gin"
	"sampleGo/controller/login"
	"sampleGo/controller/user"
	"sampleGo/handler/database/mysql"
)

func main() {
	r := gin.Default()

	// UserController 및 LoginController 인스턴스 생성
	userController := user.NewUserController()

	userGroup := r.Group("/api/user")
	{
		userGroup.GET("/:id", userController.GetUserByIDHandler)
	}

	loginGroup := r.Group("/api/login")
	{
		loginGroup.GET("/result", login.Login)
	}

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}

	dbManager := mysql.GetDBManager()
	dbManager.Close()
}
