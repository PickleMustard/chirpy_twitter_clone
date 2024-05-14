package database

import ()

type Chirp struct {
	ChirpBody string `json:"body"`
	ID        int    `json:"id"`
}
