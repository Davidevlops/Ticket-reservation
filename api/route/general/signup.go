package general

import (
	gin "github.com/gin-gonic/gin"
)

func SignUp() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	// gin.Default()
	// gin.H{"error": err.Error()}
	//  *gin.Engine
}
