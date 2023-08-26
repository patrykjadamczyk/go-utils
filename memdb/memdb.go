package memdb

import (
	"sync"
)

// Database Structure
type DB[TKey comparable, TValue any] struct {
	mu   sync.RWMutex
	data map[TKey]TValue
}

// Create Database
func Create[TKey comparable, TValue any]() *DB[TKey, TValue] {
	db := &DB[TKey, TValue]{
		data: make(map[TKey]TValue),
	}
	return db
}

// Transaction Structure
type Tx[TKey comparable, TValue any] struct {
	db       *DB[TKey, TValue]
	writable bool
}

// Lock Database
func (tx *Tx[TKey, TValue]) lock() {
	if tx.writable {
		tx.db.mu.Lock()
	} else {
		tx.db.mu.RLock()
	}
}

// Unlock Database
func (tx *Tx[TKey, TValue]) unlock() {
	if tx.writable {
		tx.db.mu.Unlock()
	} else {
		tx.db.mu.RUnlock()
	}
}

// Set Key to Value in Database
func (tx *Tx[TKey, TValue]) Set(key TKey, value TValue) {
	tx.db.data[key] = value
}

// Get Value for specified Key in Database
func (tx *Tx[TKey, TValue]) Get(key TKey) TValue {
	return tx.db.data[key]
}

// Get All Values in Database
func (tx *Tx[TKey, TValue]) GetAll() map[TKey]TValue {
	return tx.db.data
}

// View Database through Function
func (db *DB[TKey, TValue]) View(fn func(tx *Tx[TKey, TValue]) error) error {
	return db.managed(false, fn)
}

// Update Database through Function
func (db *DB[TKey, TValue]) Update(fn func(tx *Tx[TKey, TValue]) error) error {
	return db.managed(true, fn)
}

// Begin Transaction
func (db *DB[TKey, TValue]) Begin(writable bool) *Tx[TKey, TValue] {
	tx := &Tx[TKey, TValue]{
		db:       db,
		writable: writable,
	}
	tx.lock()

	return tx
}

// Managed Function for Database
func (db *DB[TKey, TValue]) managed(writable bool, fn func(tx *Tx[TKey, TValue]) error) (err error) {
	var tx *Tx[TKey, TValue]
	tx = db.Begin(writable)
	defer func() {
		if writable {
			tx.unlock()
		} else {
			tx.unlock()
		}
	}()

	_ = fn(tx)
	return
}

// Example
/*
func main() {

	db := Create()

	go db.Update(func(tx *Tx) error {
		tx.Set("mykey", "go")
		tx.Set("mykey2", "is")
		tx.Set("mykey3", "awesome")
		return nil
	})

	go db.View(func(tx *Tx) error {
		fmt.Println("value is")
		fmt.Println(tx.Get("mykey3"))
		return nil
	})

	time.Sleep(20000000000)
}
*/
