package main

import (
	"bazil.org/fuse"
	"encoding/binary"
)

type Tx struct {
	*bolt.Tx
}

func (tx *Tx) addInode(value string) Inode {
	b := tx.Bucket([]byte("inodes"))
	s, err := b.NextSequence()
	if err != nil {
		return nil
	}
	inode := Inode(s)
	b.Put(inode.Bytes(), []byte(value))
	return inode
}

func (tx *tx) addTag(name string) (Inode, error) {
	inode := addInode(name)
	b := tx.Bucket([]byte("tags"))
	err := b.Put([]byte(name), inode.Bytes())

	return inode, err
}

func (tx *Tx) findTag(name string) Inode {
	b := tx.Bucket([]byte("tags"))
	bytes := b.Get([]byte(name))
	if bytes == nil {
		return nil
	}

	return NewInode(bytes)
}
