package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sanma/domain"
)

func AddCrawlTarget(c *gin.Context) {
	c.Request.ParseForm()
	newTargetUrl := c.Request.Form["new-crawl-target"][0]

	// todo: あとで現在のユーザーを取得するように変更する
	currentUserID := "admin"

	err := domain.AddNewTarget(newTargetUrl, currentUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}
