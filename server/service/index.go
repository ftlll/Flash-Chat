package service

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// GetIndex
// @Summary ping testing
// @Schemes
// @Description ping testing
// @Tags Ping
// @Produce json
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
