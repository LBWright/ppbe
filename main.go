package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	go h.run()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/ws/:roomID", func(c *gin.Context) {
		roomID := c.Param("roomID")
		serveWs(c.Writer, c.Request, roomID)
	})

	router.Run(":" + port)
}
