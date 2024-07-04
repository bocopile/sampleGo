package main

import (
	"github.com/gin-gonic/gin"
	"sampleGo/controller/login"
	"sampleGo/controller/sample"
)

func main() {
	r := gin.Default()

	sampleGroup := r.Group("/api/sample")
	{
		sampleGroup.GET("/result", sample.Sample)
	}

	loginGroup := r.Group("/api/login")
	{
		loginGroup.GET("/result", login.Login)
	}

	r.Run("localhost:8080")
}
