package domain

import (
	"time"

	"sanma/infra"
)

type Articles struct {
	Article_id     uint64     `json:"article_id"`
	Author         string     `json:"author"`
	Article_status string     `json:"article_sattus"`
	Article_type   string     `json:"article_type"`
	Article_url    string     `json:"article_url"`
	Title          string     `json:"title"`
	Contents       string     `json:"contents"`
	Created_at     *time.Time `json:"created_at"`
	Updated_at     *time.Time `json:"updated_at"`
}

func GetAllArticles(author string) ([]*Articles, error) {
	var l []*Articles

	db := infra.Connect()
	defer db.Close()

	var err error

	if author == "_" {
		err = db.Find(&l).Error
	} else {
		err = db.Where("author = ?", author).Find(&l).Error
	}

	if err != nil {
		return nil, err
	}

	return l, err
}

func GetArticleByAID(article_id uint64) (*Articles, error) {
	db := infra.Connect()
	defer db.Close()

	searchArticle := Articles{}
	err := db.Where("article_id = ?", article_id).First(&searchArticle).Error
	if err != nil {
		return nil, err
	}

	return &searchArticle, err
}
