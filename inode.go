package main

import (
	"encoding/binary"
	"github.com/boltdb/bolt"
)

type Inode [8]byte

type InodeSet struct {
	map[Inode]int,
	empty
}

type InodeCursor struct {
	i *InodeSet,
	c *bolt.Cursor
}

func NewInode(u uint64) Inode {
	i := make(Inode)
	binary.PutUvarint(inode, uint64(u))

	return inode
}



