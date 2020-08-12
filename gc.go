package main

func LazyGarbageCollect(db *Database, key string) {
	db.DelKeyWithExpire(key)
}
