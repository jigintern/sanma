package main

import (
	"github.com/gin-gonic/gin"

	"sanma/handler"
)

func main() {
	r := gin.Default()

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/users", handler.NewUser)
		apiRouter.POST("/authentication", handler.AuthN)
		apiRouter.POST("/articles", handler.NewArticle)
		apiRouter.GET("/articles", handler.GetAllArticles)
		apiRouter.GET("/articles/:article_id", handler.GetArticlesByAID)
		apiRouter.PUT("/articles/:article_id", handler.UpdateArticle)
		apiRouter.DELETE("/articles/:article_id", handler.DeleteArticle)
		apiRouter.POST("/crawl-targets", handler.AddCrawlTarget)
	}

	r.Run(":8080")
}
