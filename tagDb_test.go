package main

import (
	"github.com/boltdb/bolt"
	"io/ioutil"
	"os"
	"testing"
)

func withDB(t testing.TB, fn func(*bolt.DB)) {
	tmp, err := ioutil.TempFile("", "yatfs-test-")
	if err != nil {
		t.Fatal(err)
	}
	dbpath := tmp.Name()
	defer func() {
		_ = os.Remove(dbpath)
	}()
	_ = tmp.Close()
	db, err := bolt.Open(dbpath, 0600, nil)
	if err != nil {
		t.Fatal(err)
	}
	fn(db)
}

func TestNewTagDB(t *testing.T) {
	withDB(t, func(db *bolt.DB) {
		_ = NewTagDB(db)
	})
}
