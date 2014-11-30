package main

import (
	"github.com/boltdb/bolt"
)

type TagDB struct {
	db *bolt.DB
}

func NewTagDB(db *bolt.DB) *TagDB {
	tagDB := TagDB{db}

	return &tagDB
}
