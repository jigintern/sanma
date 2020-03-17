package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sanma/domain"
)

func NewArticle(c *gin.Context) {}

func GetAllArticles(c *gin.Context) {
	articles, err = domain.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticlesByAID(c *gin.Context) {}

func UpdateArticle(c *gin.Context) {}

func DeleteArticle(c *gin.Context) {}
