package database

import "time"

type Token struct {
	Auth_Token      string    `json:"auth_token"`
	Refresh_Token   string    `json:"refresh_token"`
	Expiration_Date time.Time `json:"expiration_date"`
}
