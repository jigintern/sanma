package domain

import(
	"sanma/infra"
)

type Crawl_targets struct {
	Target_id	uint64	`json:"target_id"`
	Target_url	string	`json:"target_url"`
	User_id	string	`json:"user_id"`
}
