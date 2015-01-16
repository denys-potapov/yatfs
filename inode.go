package main

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
)

type Inode []byte

type InodeSet struct {
	// inodes map[Inode]int
	empty bool
}

type InodeCursor struct {
	i *InodeSet
	c *bolt.Cursor
}

func NewInode(u uint64) Inode {
	i := make(Inode, 8)
	binary.PutUvarint(i, uint64(u))

	return i
}



