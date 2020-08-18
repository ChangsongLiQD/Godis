package main

import (
	"time"
)

const (
	KeyExpireNo = -1
	KeyExpired  = 0
	KeyExpiring = 1
)

type ExpireTime int64
type Dictionary map[string]*Object
type ExpireDict map[string]ExpireTime

type Database struct {
	Dict    Dictionary
	Expires ExpireDict
}

func NewDatabase() *Database {
	return &Database{
		Dict:    map[string]*Object{},
		Expires: map[string]ExpireTime{},
	}
}

func (db *Database) GetKey(key string) *Object {
	data, _ := db.Dict[key]
	return data
}

func (db *Database) SetKey(key string, obj *Object) {
	db.Dict[key] = obj
}

func (db *Database) DelKey(key string) bool {
	_, exists := db.Dict[key]
	if exists {
		delete(db.Dict, key)
	}
	return exists
}

func (db *Database) DelKeyWithExpire(key string) bool {
	result := db.DelKey(key)
	if result {
		db.DelKeyExpire(key)
	}

	return result
}

func (db *Database) CheckExpireValid(key string) int {
	ttl, exist := db.Expires[key]
	if !exist {
		return KeyExpireNo
	}

	if ttl > ExpireTime(time.Now().UnixNano()) {
		return KeyExpiring
	}

	return KeyExpired
}

func (db *Database) ExistsKey(key string) bool {
	_, exists := db.Dict[key]
	return exists
}

func (db *Database) DelKeyExpire(key string) {
	delete(db.Expires, key)
}

func (db *Database) SetKeyExpireTimeBySeconds(key string, expire int64) bool {
	if expire > 0 {
		et := getExpireTimeBySeconds(expire)
		db.setKeyExpireTime(key, et)
		return true
	}

	return false
}

func (db *Database) setKeyExpireTime(key string, expire ExpireTime) {
	db.Expires[key] = expire
}

func (db *Database) getRandomExpireKeys(num int) ([]string, int) {
	var keys []string
	total := 0
	for key := range db.Expires {
		keys = append(keys, key)
		total++

		if total == num {
			break
		}
	}

	return keys, total
}

func getExpireTimeBySeconds(second int64) ExpireTime {
	return ExpireTime(time.Now().UnixNano() + second*1000000000)
}
