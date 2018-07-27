package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

var taskBkt = []byte("tasks")

type Tasks struct {
	ID   int
	Task string
}

var db *bolt.DB

func Init(dbPath string) {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		fmt.Printf("Error occured during db connection %s\n", err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBkt)
		return err
	})

}

func AddTask(task string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBkt)
		i, _ := b.NextSequence()
		id := int(i)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	return err
}

func ListTasks() ([]Tasks, error) {
	var tasks []Tasks
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBkt)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Tasks{
				ID:   btoi(k),
				Task: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(k int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBkt)
		return b.Delete(itob(k))
	})
}
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
