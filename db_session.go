package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"os"
)

func InitDBSession(db *bolt.DB) {
	var err error

	db, err = bolt.Open("cities.db", 0600, nil)
	if err != nil {
		fmt.Println("[DB] Couldn't connect to the database:", err.Error())
		os.Exit(1)
	}
}
