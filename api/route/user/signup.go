package user

import (
	gin "github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: hash password
	// TODO: save user to DB

	c.JSON(201, gin.H{
		"message": "user created successfully",
	})
}
