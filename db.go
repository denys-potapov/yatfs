package main

import (
	"github.com/boltdb/bolt"
)

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

// wrap bolt update transaction
func (db *DB) update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

// wrap bolt view transaction
func (db *DB) view(fn func(*Tx) error) error {
	return db.DB.View(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

func (db *DB) AddFile(tags OrderedTags, path string) error {
	return db.update(func(tx *Tx) error {
		file := tx.AddInode(path)
		for t, w := range tags {
			tag := tx.FindTag(t)
			if tag == nil {
				tag = tx.AddTag(t)
			}

			tx.AddRelation(file, tag, w)
		}

		return nil
	})
}

func (db *DB) GetFiles(tags Tags, path string) Inodes {
	t := tagsToInodes
}