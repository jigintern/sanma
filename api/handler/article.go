package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"sanma/domain"
)

func NewArticle(c *gin.Context) {}

func GetAllArticles(c *gin.Context) {
	user_id := c.DefaultQuery("user_id", "_")

	if user_id == "_" {
		articles, err := domain.GetAllArticles()
	} else {
		articles, err := domain.GetArticlesByUID(user_id)
	}

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
