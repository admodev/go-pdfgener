package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeServer() error {
	r := gin.Default();
	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Captura el error retornado por r.Run y lo pasa al caller
	return r.Run() // listen and serve by default on 0.0.0.0:8080 (for windows "localhost:8080")
}
