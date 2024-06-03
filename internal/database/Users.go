package database

import ()

type User struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	EncryptedHash []byte `json:"encrypted_password"`
	Id            int    `json:"id"`
}
