package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetCurrentTime first demo function
func GetCurrentTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"current_time": time.Now().Format("2006-01-02 15:04:05"),
	})
	return
}
