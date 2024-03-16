package user

import (
	"fmt"

	"gitlab.com/arnaud-web/neli-backoffice/api"
	"gitlab.com/arnaud-web/neli-backoffice/config"
)

type user struct {
	UserID    int64  `json:"userId"`
	Email     string `json:"email"`
	Lastname  string `json:"lastName"`
	Firstname string `json:"firstName"`
	Role      string `json:"profile"`
}

func Info(token string) (*user, error) {
	u := new(user)
	url := fmt.Sprintf("%s/users/me", *config.ApiURL)

	if err := api.Get(url, token, u); err != nil {
		return u, err
	}

	return u, nil
}
