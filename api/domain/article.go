package domain

import (
	"sanma/infra"
	"time"
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
