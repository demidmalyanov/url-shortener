package main

import (
	"fmt"
	"log"

	"github.com/demidmalyanov/url-shortener/database"
	"github.com/demidmalyanov/url-shortener/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "URL shortener , please create link",
		})
	})

	r.POST("/shorten-url", func(c *gin.Context) {
		handlers.CreateShortURL(c, db)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleURLRedirect(c, db)
	})

	err = r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
