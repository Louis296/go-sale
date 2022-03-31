package handler

import "github.com/gin-gonic/gin"

func (receiver Handler) Test20220101(c *gin.Context) {
	c.JSON(200, gin.H{"message": "test"})
}
