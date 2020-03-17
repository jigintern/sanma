package domain

import (
	"sanma/infra"
)

type Crawl_targets struct {
	Target_id  uint64 `json:"target_id"`
	Target_url string `json:"target_url"`
	User_id    string `json:"user_id"`
}

func AddNewTarget(newtargeturl string, user_id string) error {
	db := infra.Connect()
	defer db.Close()

	newTarget := Crawl_targets{}
	newTarget.Target_url = newtargeturl
	newTarget.User_id = user_id

	return db.Create(&newTarget).Error
}
