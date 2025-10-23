package single

import "sync"

type DB struct {
	Url string
}

func NewDB(url string) *DB {
	return &DB{url}
}

var (
	db   *DB
	once sync.Once
)

func getDB() *DB {
	if db == nil {
		once.Do(func() { // Avoid creating multiple instance in concurrent scenarios
			db = NewDB("http://www.baidu.com")
		})
	}
	return db
}
