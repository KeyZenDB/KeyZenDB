package internal

import (
	"sync"
	"time"
)

type Db struct {
	store sync.Map
}

func NewDb() *Db {
	return &Db{}
}

func (db *Db) Set(key string, value any, ttl ...time.Duration) any {
	db.store.Store(key, value)
	if len(ttl) > 0 && ttl[0] > 0 {
		// if have ttl after this time remove this key
		go func() {
			time.Sleep(ttl[0])
			db.Del(key)
		}()
	}
	return "OK"
}

func (db *Db) Get(key string) any {
	if value, ok := db.store.Load(key); ok {
		return value
	}
	return nil
}

func (db *Db) Del(key string) any {
	db.store.Delete(key)
	return "OK"
}

func (db *Db) Clear() any {
	db.store.Clear()
	return "OK"
}

func (db *Db) Keys() []string {
	keys := []string{}
	db.store.Range(func(k, _ any) bool {
		keys = append(keys, k.(string))
		return true
	})
	return keys
}

func (db *Db) Exists(key string) bool {
	_, ok := db.store.Load(key)
	return ok
}

func (db *Db) Incr(key string) int {
	value, _ := db.store.Load(key)
	num, ok := value.(int)
	if !ok {
		num = 0
	}
	num++
	db.store.Store(key, num)
	return num
}

func (db *Db) Decr(key string) int {
	value, _ := db.store.Load(key)
	num, ok := value.(int)
	if !ok {
		num = 0
	}
	num--
	db.store.Store(key, num)
	return num
}

func (db *Db) Len() int {
	count := 0
	db.store.Range(func(_, _ any) bool {
		count++
		return true
	})
	return count
}

func (db *Db) Ping() any {
	return "PONG"
}
