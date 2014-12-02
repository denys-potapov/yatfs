package main

import (
	"bazil.org/fuse"
	"encoding/binary"
	"github.com/boltdb/bolt"
)

type Inode fuse.NodeID

type DB struct {
	*bolt.DB
}

type Tx struct {
	*bolt.Tx
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

func (inode Inode) Bytes() []byte {
	bytes := make([]byte, 8, 8)
	binary.PutUvarint(bytes, uint64(inode))
	return bytes
}

// wrap bolt transaction
func (db *DB) Update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

func (db *DB) addInode(value string) (Inode, error) {
	var inode Inode
	err := db.Update(func(tx *Tx) error {
		b := tx.Bucket([]byte("inodes"))
		s, _ := b.NextSequence()
		inode := Inode(s)
		b.Put(inode.Bytes(), []byte(value))
		return nil
	})

	return inode, err
}
