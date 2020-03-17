package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"sanma/domain"
)

func NewArticle(c *gin.Context) {}

func GetAllArticles(c *gin.Context) {
	user_id := c.DefaultQuery("user_id", "_")
	articles, err := domain.GetAllArticles(user_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, articles)
}

func GetArticlesByAID(c *gin.Context) {
	article_id, err := strconv.ParseUint(c.Param("article_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	article, err := domain.GetArticleByAID(article_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, article)
}

func UpdateArticle(c *gin.Context) {}

func DeleteArticle(c *gin.Context) {}
