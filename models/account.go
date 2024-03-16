package models

type Account struct {
	UserID    int64  `json:"userId,omitempty"`
	Email     string `json:"email,omitempty"`
	Lastname  string `json:"lastName,omitempty"`
	Firstname string `json:"firstName,omitempty"`
	Role      string `json:"profile,omitempty"`
}
