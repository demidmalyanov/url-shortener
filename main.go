package main

import (
	"context"
	"fmt"
	"log"

	"github.com/demidmalyanov/url-shortener/handlers"
	"github.com/demidmalyanov/url-shortener/storage/sqlite"
	"github.com/gin-gonic/gin"
)

const (
	storageSqlitePath = "data/tokens.db"
)

func main() {
	s, err := sqlite.New(storageSqlitePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "URL shortener , please create link",
		})
	})

	r.POST("/shorten-url", func(c *gin.Context) {
		handlers.CreateShortURL(c,s)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleURLRedirect(c,s)
	})

	err = r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
