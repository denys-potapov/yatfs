package main

import (
	"github.com/boltdb/bolt"
	"io/ioutil"
	"os"
	"testing"
)

func withDB(t *testing.T, fn func(*bolt.DB)) {
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

func TestInitDB(t *testing.T) {
	withDB(t, func(db *bolt.DB) {
		_ = InitDB(db)

		db.View(func(tx *bolt.Tx) error {
			for _, name := range []string{"tags", "inodes", "relations"} {
				if tx.Bucket([]byte(name)) == nil {
					t.Errorf("no %s bucket", name)
				}
			}
			return nil
		})
	})
}
