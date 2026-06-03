package sbicmn

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Login(c *gin.Context) {
	c.String(200, "Login")
}

func Sub(c *gin.Context) {
	c.String(200, "Sub")
}

func TestLoggerToFile(t *testing.T) {

	router := gin.Default()
	router.Use(LoggerToFile("test.log"))

	// ¼òµ¥µÄ×éÂ·ÓÉ: v1
	v1 := router.Group("/v1", gin.Logger())
	{
		v1.GET("/login", Login)
		v1.GET("/submit", Sub)
	}

	// ¼òµ¥µÄ×éÂ·ÓÉ: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/log", Login)
		v2.GET("/sub", Sub)
	}

	router.Run(":8080")

}
