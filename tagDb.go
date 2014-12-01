package main

import (
	"github.com/boltdb/bolt"
	"bazil.org/fuse"
)

type TagDB struct {
	db *bolt.DB
}

type inode fuse.NodeID

func NewTagDB(db *bolt.DB) *TagDB {

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
	t := TagDB{db}

	return &t
}

func getTagId() inode {

}

func (t *TagDB) AddFile(tags []string, path string) (inode, error) {

}
