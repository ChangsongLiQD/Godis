package main

type Dictionary map[string]*Object

type Database struct {
	Dict   Dictionary
	Expire Dictionary
}

func NewDatabase() *Database {
	return &Database{
		map[string]*Object{},
		map[string]*Object{},
	}
}

func (db *Database) SetKey(key string, obj *Object) {
	db.Dict[key] = obj
}

func (db *Database) ExistsKey(key string) bool {
	_, exists := db.Dict[key]
	return exists
}
