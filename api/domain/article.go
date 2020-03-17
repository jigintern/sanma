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

func GetAllArticles() ([]*Articles, error) {
	var l []*Articles

	db := infra.Connect()
	defer db.Close()

	err := db.Find(&l).Error
	if err != nil {
		return nil, err
	}

	return l, err
}
