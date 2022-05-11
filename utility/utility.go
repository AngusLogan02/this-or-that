package utility

import (
	"log"
	"strconv"
	"time"

	bolt "go.etcd.io/bbolt"
)

func Init(file string) {
	db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Video_game"))
		return err
	})
	db.Close()
}

func Open(file string) *bolt.DB {
	db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	return db
}

func Set(db *bolt.DB, bucket, key, value string) {
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func Get(db *bolt.DB, bucket, key string) string {
	var s []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		s = b.Get([]byte(key))
		if s == nil {
			s = []byte("0")
		}
		return nil
	})

	return string(s)
}

func Del(db *bolt.DB, bucket, key string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Delete([]byte(key))
		return nil
	})
}

func Increment(key string) {
	db := Open("my.db")

	currentCount := Get(db, "Video_game", key)

	currentCountInt, err := strconv.Atoi(currentCount)
	if err != nil {
		log.Fatal("Something's up with the database. Error converting to int.", err)
	}

	var newCount int
	newCount = currentCountInt + 1
	newCountStr := strconv.Itoa(newCount)

	Set(db, "Video_game", key, newCountStr)

	db.Close()
}

func Iterate(bucket string) (keys []string, values []string) {
	keys = make([]string, 0)
	values = make([]string, 0)

	db := Open("my.db")
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()
		i := 0
		for k, v := c.First(); k != nil && i < 50; k, v = c.Next() {
			keys = append(keys, string(k))
			values = append(values, string(v))
			i++
		}

		return nil
	})
	db.Close()
	return keys, values
}
