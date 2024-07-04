package sample

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "helloworld",
	})

}
