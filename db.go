package main

import (
	"bazil.org/fuse"
	"encoding/binary"
	"github.com/boltdb/bolt"
)

type Inode fuse.NodeID

func (inode Inode) Bytes() []byte {
	bytes := make([]byte, 8, 8)
	binary.PutUvarint(bytes, uint64(inode))
	return bytes
}

func NewInode(bytes []byte) Inode {
	inode, _ = binary.Uvarint(bytes)
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
func (db *DB) Update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}
