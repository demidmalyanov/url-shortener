package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/demidmalyanov/url-shortener/database"
	"github.com/demidmalyanov/url-shortener/models"
	"github.com/demidmalyanov/url-shortener/shortener"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	Url string `json:"url" binding:"required"`
}

const host = "http://localhost:8080/"

func CreateShortURL(c *gin.Context, db *database.TokenDB) {
	var createRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	urlToken := shortener.GenerateTokenForUrl(createRequest.Url)

	// save in db
	_, err := db.Insert(models.Token{
		Token: urlToken,
		Url:   createRequest.Url,
	})

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + urlToken,
	})
}

func HandleURLRedirect(c *gin.Context, db *database.TokenDB) {

	shortUrl := c.Param("shortUrl")
	// save in db
	fmt.Println(shortUrl)

	initialUrl, err := db.GetUrlByToken(shortUrl)

	fmt.Println("ggg", initialUrl)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(302, initialUrl)

}
