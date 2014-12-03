package main

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
)

type Inode []byte

func NewInode(i uint64) Inode {
	inode := make([]byte, 8)
	binary.PutUvarint(inode, uint64(i))

	return inode
}

type DB struct {
	*bolt.DB
}

func InitDB(db *bolt.DB) *DB {
	db.Update(func(tx *bolt.Tx) error {
		buckets := []string{"tags", "inodes", "relations"}
		for _, name := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
		}
		return nil
	})

	return &DB{db}
}

// wrap bolt transaction
func (db *DB) update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}
