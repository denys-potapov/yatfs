package main

import (
	"github.com/boltdb/bolt"
)

type Tx struct {
	*bolt.Tx
}

func (tx *Tx) addInode(value string) Inode {
	b := tx.Bucket([]byte("inodes"))
	i, err := b.NextSequence()
	if err != nil {
		return nil
	}
	inode := NewInode(i)
	b.Put(inode, []byte(value))

	return inode
}

func (tx *Tx) AddTag(name string) Inode {
	inode := tx.addInode(name)
	b := tx.Bucket([]byte("tags"))
	b.Put([]byte(name), inode)

	return inode
}

func (tx *Tx) FindTag(name string) Inode {
	b := tx.Bucket([]byte("tags"))
	inode := b.Get([]byte(name))

	return inode
}

func (tx *Tx) addRelation(i1 Inode, i2 Inode, weight byte) error {
	b := tx.Bucket([]byte("relations"))
	item, err := b.CreateBucketIfNotExists(i1)
	if err != nil {
		return err
	}

	return item.Put(i2, []byte{weight})
}

func (tx *Tx) AddRelation(file Inode, tag Inode, weight byte) error {
	err := tx.addRelation(file, tag, weight)
	if err != nil {
		return err
	}
	err = tx.addRelation(tag, file, weight)
	return err
}
