package main

import (
	"bytes"
	"github.com/boltdb/bolt"
	"testing"
)

func withTx(t *testing.T, fn func(*Tx)) {
	withDB(t, func(db *bolt.DB) {
		d := InitDB(db)
		d.update(func(tx *Tx) error {
			fn(tx)
			return nil
		})
	})
}

func TestAddTag(t *testing.T) {
	withTx(t, func(tx *Tx) {
		inode := tx.AddTag("tag")

		i := tx.Bucket([]byte("inodes"))
		if g, e := string(i.Get(inode)), "tag"; g != e {
			t.Error("error adding tag inode")
		}

		b := tx.Bucket([]byte("tags"))
		if g, e := b.Get([]byte("tag")), inode; !bytes.Equal(g, e) {
			t.Errorf("error adding tag %v %v", g, e)
		}
	})
}

func TestFindTag(t *testing.T) {
	withTx(t, func(tx *Tx) {
		inode := NewInode(1)
		_ = tx.Bucket([]byte("tags")).Put([]byte("tag"), inode)

		if g, e := tx.FindTag("tag"), inode; !bytes.Equal(g, e) {
			t.Error("couldnt find tag")
		}
	})
}

func TestAddRelation(t *testing.T) {
	withTx(t, func(tx *Tx) {
		file, tag, weight := NewInode(1), NewInode(2), byte(5)

		_ = tx.AddRelation(file, tag, weight)

		b := tx.Bucket([]byte("relations"))

		tags := b.Bucket(file)
		if g, e := tags.Get(tag), []byte{weight}; !bytes.Equal(g, e) {
			t.Errorf("error adding file tags")
		}

		files := b.Bucket(tag)
		if g, e := files.Get(file), []byte{weight}; !bytes.Equal(g, e) {
			t.Errorf("error adding tags file")
		}
	})
}
