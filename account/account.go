package account

import (
	"fmt"

	"gitlab.com/arnaud-web/neli-backoffice/api"
	"gitlab.com/arnaud-web/neli-backoffice/config"
)

type Account struct {
	UserID    int64  `json:"userId"`
	Email     string `json:"email"`
	Lastname  string `json:"lastName"`
	Firstname string `json:"firstName"`
	Role      string `json:"profile"`
}

func Info(token string) (*Account, error) {
	u := new(Account)
	url := fmt.Sprintf("%s/users/me", *config.ApiURL)

	if err := api.Get(url, token, u); err != nil {
		return u, err
	}

	return u, nil
}
