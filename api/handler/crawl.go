package handler

import(
	"github.com/gin-gonic/gin"

	"sanma/domain"
)

func AddCrawlTarget(c *gin.Context) {
	c.Request.ParseForm()
	newTarget := c.Request.Form["new-crawl-target"]

	// todo: あとで現在のユーザーを取得するように変更する
	currentUser := "admin"

	domain.
}
