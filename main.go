package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	router := gin.New()

	router.GET("/ws/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		serveWs(c.Writer, c.Request, roomID)
	})

	router.Run("0.0.0.0:8082")
}
