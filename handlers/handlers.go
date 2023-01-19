package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/demidmalyanov/url-shortener/shortener"
	"github.com/demidmalyanov/url-shortener/storage"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	Url string `json:"url" binding:"required"`
}

const host = "http://localhost:8080/"

func CreateShortURL(c *gin.Context, s storage.Storage) {
	var createRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlToken := shortener.GenerateTokenForUrl(createRequest.Url)

	token := storage.Token{
		Url:   createRequest.Url,
		Token: urlToken,
	}

	// save in db
	err := s.Save(context.Background(), &token)
	if err != nil {
		log.Fatal("can`t save in db:", err)
	}

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + urlToken,
	})
}

func HandleURLRedirect(c *gin.Context, s storage.Storage) {

	shortUrl := c.Param("shortUrl")
	// save in db
	fmt.Println(shortUrl)

	t, err := s.Get(context.Background(), shortUrl)

	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(302, t.Url)

}
