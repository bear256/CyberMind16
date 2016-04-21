package database

import (
	"log"

	"github.com/boltdb/bolt"
)

var Db *bolt.DB

func init() {
	db, err := bolt.Open("fia.db", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	Db = db
}
