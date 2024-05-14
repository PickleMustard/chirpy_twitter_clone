package database

import (
	"encoding/json"
	"errors"
	//uuid "github.com/google/uuid"
	"io/fs"
	"log"
	"os"
	"sync"
)

type DB struct {
	path          string
	mux           *sync.RWMutex
	stored_values DBStructure
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
	Users  map[int]User  `json:"users"`
}

func NewDB(path string) (*DB, error) {
	const file_header = ""
	_, exist := os.ReadFile(path)
	if exist != fs.ErrNotExist {
		os.Remove(path)
	}

	_database := DB{
		path:          path,
		mux:           &sync.RWMutex{},
		stored_values: DBStructure{Chirps: make(map[int]Chirp), Users: make(map[int]User)},
	}

	return &_database, nil
}

func (db *DB) CreateChirp(body string) (Chirp, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	//Chirp should be validated by endpoint before being written to DB
	//If I'm here, then it has been validated
	chirp := Chirp{
		ChirpBody: body,
		ID:        len(db.stored_values.Chirps) + 1,
	}

	db.stored_values.Chirps[chirp.ID] = chirp

	db_err := db.writeDB(db.stored_values)

	if db_err != nil {
		return Chirp{}, db_err
	}

	return chirp, nil
}

func (db *DB) CreateUser(email string) (User, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	user := User{
		Email: email,
		Id:    len(db.stored_values.Users) + 1,
	}

	db.stored_values.Users[user.Id] = user

	db_err := db.writeDB(db.stored_values)

	if db_err != nil {
		return User{}, db_err
	}

	return user, nil
}

func (db *DB) GetChirps() ([]Chirp, error) {
	struc, db_err := db.loadDB()
	array := make([]Chirp, 0)
	if db_err != nil {
		log.Println("Error reading from database")
		return nil, db_err
	}

	for _, text := range struc.Chirps {
		log.Println(text)
		array = append(array, text)
	}

	return array, nil
}

func (db *DB) GetSpecificChirps(id int) (Chirp, error) {
	struc, db_err := db.loadDB()
	if db_err != nil {
		log.Println("Error reading from database")
		return Chirp{}, db_err
	}

	found_value, err := struc.Chirps[id]
	log.Println(found_value)
	if !err {
		log.Printf("Could not find chirp")
		return Chirp{}, errors.New("Could not find chirp")
	}
	return found_value, nil
}

func (db *DB) ensureDB() error { return nil }

func (db *DB) loadDB() (DBStructure, error) {
	var chirps = DBStructure{}
	log.Printf("Reading file: %s", db.path)

	file_data, read_err := os.ReadFile(db.path)

	if read_err != nil {
		log.Println("Error Reading data")
		return DBStructure{}, read_err
	}

	json_err := json.Unmarshal(file_data, &chirps)
	if json_err != nil {
		log.Printf("Error reading from json: %s", json_err)
		return DBStructure{}, json_err
	}

	return chirps, nil
}

func (db *DB) writeDB(dbStructure DBStructure) error {
	marshalled_chirp, err := json.Marshal(dbStructure)

	if err != nil {
		log.Println("Error marshalling database structure")
		return err
	}

	file_err := os.WriteFile(db.path, marshalled_chirp, 0666)
	if file_err != nil {
		log.Println("Error writing json to file")
		return file_err
	}
	return nil
}
