package database

import ()

type Chirp struct {
	ChirpBody string `json:"body"`
	ID        int    `json:"id"`
	Author_ID int    `json:"author_id"`
}
